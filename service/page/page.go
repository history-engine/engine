package page

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
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
		logger.Zap().Error("save page error", zap.Error(err), zap.String("sql", sql), zap.Any("page", page))
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
