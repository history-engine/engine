package page

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
)

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
