package page

import (
	"context"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/service/readability"
	"time"
)

// SavePage 保存页面
func SavePage(ctx context.Context, page *model.Page) (int64, error) {
	if page.Version == 0 {
		page.Version = NextVersion(ctx, page.UniqueId)
	}

	if page.UpdatedAt.IsZero() {
		page.UpdatedAt = time.Now()
	}

	readability.Parser()
	x := db.GetEngine()
	sql := "insert into page set " +
		"user_id=:user_id, unique_id=:unique_id, version=:version, title=:title, " +
		"url=:url, full_path=:full_path, full_size=:full_size, lite_path=:lite_path, " +
		"lite_size=:lite_size, indexed_at=:indexed_at, updated_at=:updated_at"
	res, err := x.NamedExecContext(ctx, sql, page)
	if err != nil {
		panic(err)
	}

	page.Id, err = res.LastInsertId()
	if err != nil {
		panic(err)
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
