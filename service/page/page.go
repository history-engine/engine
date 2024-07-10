package page

import (
	"context"
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

func Search(ctx context.Context, search model.SearchPage) (maps map[string][]model.PageSearch, total int, err error) {
	zincSearch, err := zincsearch.EsSearch(search)
	if err != nil {
		panic(err)
	}

	if zincSearch.Hits.Total.Value == 0 {
		return
	}

	docs := make(map[string]model.ZincDocument, 0)
	ids := make([]string, zincSearch.Hits.Total.Value)
	for k, v := range zincSearch.Hits.Hits {
		ids[k] = v.ID
		if source, ok := v.Source.(map[string]interface{}); ok {
			docs[v.ID] = model.ZincDocument{
				Id:      v.ID,
				Title:   source["title"].(string),
				Content: source["content"].(string),
				Url:     source["url"].(string),
				Size:    int(source["size"].(float64)),
			}
		}
	}

	pages, err := BatchGetPage(ctx, ids)
	maps = map[string][]model.PageSearch{}
	for _, v := range pages {
		v.FullPath = setting.Web.Domain + "/page/view" + v.FullPath
		if _, ok := maps[v.UniqueId]; !ok {
			maps[v.UniqueId] = []model.PageSearch{}
		}
		maps[v.UniqueId] = append(maps[v.UniqueId], model.PageSearch{
			Avatar:  "https://avatars.akamai.steamstatic.com/6a9ae9c069cd4fff8bf954938727730cdb0fe27b.jpg",
			Title:   docs[v.UniqueId].Title,
			Content: docs[v.UniqueId].Content,
			Url:     docs[v.UniqueId].Url,
			Size:    docs[v.UniqueId].Size,
			Preview: v.FullPath,
		})
	}

	return maps, zincSearch.Hits.Total.Value, nil
}
