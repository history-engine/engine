package page

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	entPage "history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/search"
	"history-engine/engine/setting"
	"os"
	"time"
)

func Versions(ctx context.Context, userId int64, uniqueId string, page, limit int) (int, []model.SearchResultPage, error) {
	x := db.GetEngine()

	pageQuery := x.Page.Query().Where(entPage.UserID(userId), entPage.UniqueID(uniqueId))

	total, err := pageQuery.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	source, err := pageQuery.
		Order(ent.Desc(entPage.FieldID)).
		Offset((page - 1) * limit).
		Limit(limit).
		All(ctx)

	pages := make([]model.SearchResultPage, 0)
	for _, item := range source {
		pages = append(pages, EntPage2SearchResultPage(ctx, item))
	}

	return total, pages, err
}

// NextVersion 获取下一个版本号
func NextVersion(ctx context.Context, uniqueId string) (int, time.Time) {
	x := db.GetEngine()
	page, err := x.Page.Query().
		Select(entPage.FieldID, entPage.FieldVersion, entPage.FieldCreatedAt).
		Where(entPage.UniqueID(uniqueId)).
		Order(entPage.ByVersion(sql.OrderDesc())).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return 1, time.Time{}
		}

		logger.Zap().Warn("get max version error", zap.String("unique_id", uniqueId), zap.Error(err))
		return 0, time.Now()
	}

	return page.Version + 1, page.CreatedAt
}

// CleanHistory 清除历史版本、HTML文件、ZincSearch索引
func CleanHistory(ctx context.Context, userId int64, uniqueId string, version int) error {
	diff := version - setting.SingleFile.MaxVersion
	if diff <= 0 {
		return nil
	}

	x := db.GetEngine()
	var vs []struct {
		Id      int64
		Version int
		Path    string
	}
	if err := x.Page.Query().
		Where(entPage.UserID(userId), entPage.UniqueID(uniqueId), entPage.VersionLTE(diff)).
		Select(entPage.FieldID, entPage.FieldVersion, entPage.FieldPath).
		Scan(ctx, &vs); err != nil {
		return err
	}

	var err error
	for _, v := range vs {
		docId := fmt.Sprintf("%s%d", uniqueId, version)
		if err := search.Engine().DelDocument(ctx, userId, docId); err != nil {
			err = errors.Join(err)
			continue
		}

		if err := os.Remove(setting.SingleFile.HtmlPath + v.Path); err != nil {
			err = errors.Join(err)
			continue
		}

		if err := x.Page.DeleteOneID(v.Id).Exec(ctx); err != nil {
			logger.Zap().Error("delete page err", zap.Int64("id", v.Id))
			err = errors.Join(err)
			continue
		}
	}

	return err
}
