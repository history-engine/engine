package page

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/zincsearch"
	"history-engine/engine/setting"
	"sync"
	"time"
)

var pageLock = sync.Mutex{}

// SavePage 保存页面
func SavePage(ctx context.Context, page *model.Page) (int64, error) {
	if page.Version == 0 {
		page.Version = NextVersion(ctx, page.UniqueId)
	}

	if page.UpdatedAt.IsZero() {
		page.UpdatedAt = time.Now()
	}

	pageLock.Lock()
	defer pageLock.Unlock()
	x := db.GetEngine()
	sql := "insert into page set " +
		"user_id=:user_id, unique_id=:unique_id, version=:version, title=:title, " +
		"url=:url, full_path=:full_path, full_size=:full_size, lite_path=:lite_path, " +
		"lite_size=:lite_size, indexed_at=:indexed_at, updated_at=:updated_at"
	res, err := x.NamedExecContext(ctx, sql, page)
	if err != nil {
		version := NextVersion(ctx, page.UniqueId)
		logger.Zap().Error("save page error", zap.Error(err), zap.String("sql", sql), zap.Any("page", page), zap.Int("version", version))
		return 0, err
	}

	page.Id, err = res.LastInsertId()
	if err != nil {
		logger.Zap().Error("get last insert id error", zap.Error(err), zap.String("sql", sql), zap.Any("page", page))
		return 0, err
	}

	// 清除历史版本
	go func() {
		err := CleanHistory(context.Background(), page.UniqueId)
		if err != nil {
			panic(err)
		}
	}()

	return page.Id, nil
}

// TODO 少获取几个字段
func BatchGetPage(ctx context.Context, uniqueId []string) ([]model.Page, error) {
	if len(uniqueId) == 0 {
		return nil, nil
	}

	x := db.GetEngine()
	var pages []model.Page
	query, args, err := sqlx.In("select * from page where unique_id in (?) order by created_at desc", uniqueId)
	if err != nil {
		panic(err)
	}

	err = x.SelectContext(ctx, &pages, query, args...)
	if err != nil {
		panic(err)
	}

	return pages, nil
}

func Page(ctx context.Context, start, rows int) ([]model.Page, error) {
	x := db.GetEngine()
	var list []model.Page
	err := x.SelectContext(ctx, &list, "select * from page order by created_at desc limit ?, ?", start, rows)
	if err != nil {
		panic(err)
	}

	return list, err
}

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
	docMap := make(map[string]model.Page, 0)
	for _, v := range pages {
		docMap[fmt.Sprintf("%s%d", v.UniqueId, v.Version)] = v
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
			Title:   page.Title,
			Url:     page.Url,
			Size:    page.FullSize,
			Preview: setting.Web.Domain + "/page/view" + page.FullPath,
			Version: page.Version,
		}

		if source, ok := v.Source.(map[string]interface{}); ok {
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
