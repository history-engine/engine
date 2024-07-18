package page

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/zincsearch"
	"history-engine/engine/setting"
	"os"
	"time"
)

// NextVersion 获取下一个版本号
func NextVersion(ctx context.Context, uniqueId string) (int, time.Time) {
	x := db.GetEngine()
	page, err := x.Page.Query().
		Select(page.FieldID, page.FieldVersion, page.FieldCreatedAt).
		Where(page.UniqueID(uniqueId)).
		Order(page.ByVersion(sql.OrderDesc())).
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
func CleanHistory(ctx context.Context, userId int, uniqueId string, version int) {
	diff := version - setting.SingleFile.MaxVersion
	if diff <= 0 {
		return
	}

	x := db.GetEngine()
	var vs []struct {
		Id      int64
		Version int
		Path    string
	}
	if err := x.Page.Query().
		Where(page.UserID(userId), page.UniqueID(uniqueId), page.VersionLTE(diff)).
		Select(page.FieldID, page.FieldVersion, page.FieldPath).
		Scan(ctx, &vs); err != nil {
		return
	}

	for _, v := range vs {
		if err := x.Page.DeleteOneID(v.Id).Exec(ctx); err != nil {
			logger.Zap().Error("delete page err", zap.Int64("id", v.Id))
			continue
		}

		if err := os.Remove(setting.SingleFile.HtmlPath + v.Path); err != nil {
			logger.Zap().Error("delete html file err", zap.String("path", v.Path))
			continue
		}

		if err := zincsearch.DelDocument(userId, uniqueId, v.Version); err != nil {
			logger.Zap().Error(
				"delete zinc search doc err",
				zap.String("unique_id", uniqueId),
				zap.Int("version", v.Version),
			)
		}
	}

	return
}
