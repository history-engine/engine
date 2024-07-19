package page

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/readability"
	"history-engine/engine/service/zincsearch"
	"history-engine/engine/setting"
	"sync"
	"time"
)

// ParserPage 调用readability分析HTML文件，添加到ZincSearch、保存数据库
func ParserPage(ctx context.Context, uniqueId string) {
	pages, err := BatchGetPage(ctx, []string{uniqueId})
	if err != nil {
		logger.Zap().Warn("parse page err", zap.Error(err), zap.String("unique_id", uniqueId))
		return
	}

	x := db.GetEngine()
	for _, v := range pages {
		if !v.IndexedAt.IsZero() {
			continue
		}

		article, err := readability.Parser().Parse(setting.SingleFile.HtmlPath + v.Path)
		if err != nil {
			logger.Zap().Warn("parse page err", zap.Error(err), zap.String("unique_id", uniqueId))
			continue
		}

		zincId := fmt.Sprintf("%s%d", uniqueId, v.Version)
		zincDoc := &model.ZincDocument{
			Url:     v.URL,
			Title:   article.Title,
			Excerpt: article.Excerpt,
			Content: article.TextContent,
		}
		if err = zincsearch.PutDocument(v.UserID, zincId, zincDoc); err != nil {
			logger.Zap().Warn("put zinc doc err", zap.Error(err), zap.String("unique_id", uniqueId))
			continue
		}

		_, err = x.Page.Update().SetTitle(article.Title).SetIndexedAt(time.Now()).Where(page.ID(v.ID)).Save(ctx)
		if err != nil {
			logger.Zap().Warn("update page err", zap.Error(err), zap.String("unique_id", uniqueId))
			continue
		}
	}
}

var pageLock = sync.Mutex{}

// SavePage 保存页面
func SavePage(ctx context.Context, page *ent.Page) (*ent.Page, error) {
	pageLock.Lock()
	defer pageLock.Unlock()

	x := db.GetEngine()

	// todo 这里有点繁琐和多余
	return x.Page.Create().
		SetUserID(page.UserID).
		SetUniqueID(page.UniqueID).
		SetVersion(page.Version).
		SetTitle(page.Title).
		SetURL(page.URL).
		SetPath(page.Path).
		SetSize(page.Size).
		Save(ctx)
}

// BatchGetPage TODO 少获取几个字段
func BatchGetPage(ctx context.Context, uniqueId []string) ([]*ent.Page, error) {
	if len(uniqueId) == 0 {
		return nil, nil
	}

	x := db.GetEngine()

	pages, err := x.Page.Query().
		Where(page.UniqueIDIn(uniqueId...)).
		Order(page.ByCreatedAt(sql.OrderDesc())).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return pages, nil
}

// Page 分页获取页面
func Page(ctx context.Context, start, rows int) ([]*ent.Page, error) {
	x := db.GetEngine()
	return x.Page.Query().Offset(start).Limit(rows).All(ctx)
}

// Search 页面搜索
func Search(ctx context.Context, userId int64, search model.SearchPage) (pageSearchs []model.PageSearch, total int, err error) {
	zincSearch, err := zincsearch.EsSearch(userId, search)
	if err != nil {
		panic(err)
	}

	if zincSearch.Hits.Total.Value == 0 {
		return
	}

	// 提取页面id
	docIdList := make([]string, 0)
	for _, v := range zincSearch.Hits.Hits {
		docIdList = append(docIdList, v.ID[0:32])
	}

	// 从数据获取页面信息
	pages, err := BatchGetPage(ctx, docIdList)
	docMap := make(map[string]*ent.Page, 0)
	for _, v := range pages {
		docMap[fmt.Sprintf("%s%d", v.UniqueID, v.Version)] = v
	}

	// 提取Zinc结果，部分数据从数据库补全
	pageSearchs = make([]model.PageSearch, 0)
	for _, v := range zincSearch.Hits.Hits {
		page, ok := docMap[v.ID]
		if !ok {
			continue
		}

		pageSearch := model.PageSearch{
			Avatar:  "https://avatars.akamai.steamstatic.com/6a9ae9c069cd4fff8bf954938727730cdb0fe27b.jpg",
			Url:     page.URL,
			Size:    page.Size,
			Preview: setting.Web.Domain + "/page/view" + page.Path,
			Version: page.Version,
		}

		if source, ok := v.Source.(map[string]interface{}); ok {
			if val, ok := source["title"].(string); ok {
				pageSearch.Title = val
			}
			if val, ok := source["excerpt"].(string); ok {
				pageSearch.Excerpt = val
			}
			if val, ok := source["content"].(string); ok {
				pageSearch.Content = val
			}
			pageSearchs = append(pageSearchs, pageSearch)
		}
	}

	return pageSearchs, zincSearch.Hits.Total.Value, nil
}
