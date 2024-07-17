package page

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"log"
	"os"
	"time"
)

// NextVersion 获取下一个版本号
func NextVersion(ctx context.Context, uniqueId string) (int, time.Time) {
	page := model.Page{}

	x := db.GetEngine()
	query := "select id, version, created_at from page where unique_id=? order by version desc limit 1"
	err := x.GetContext(ctx, &page, query, uniqueId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Zap().Warn("get max version error", zap.String("sql", query), zap.String("unique_id", uniqueId), zap.Error(err))
	}

	return page.Version + 1, page.CreatedAt
}

// CleanHistory 清除历史版本
// TODO 清理对应的文件
func CleanHistory(ctx context.Context, uniqueId string) error {
	x := db.GetEngine()

	// 获取最大最小版本号
	v := struct {
		Min int
		Max int
	}{
		Min: 0,
		Max: 0,
	}

	sql := "select min(version) as min, max(version) as max from page where unique_id=?"
	err := x.GetContext(ctx, &v, sql, uniqueId)
	if err != nil {
		logger.Zap().Error("check min and max version error", zap.String("sql", sql), zap.String("unique_id", uniqueId), zap.Error(err))
		return err
	}

	if v.Max-v.Min < setting.SingleFile.MaxVersion {
		return nil
	}

	// todo 改成全部查出来再删除比较好因为还要删除文件和索引
	sql = "delete from page where unique_id=? and version < ?"
	res, err := x.ExecContext(ctx, sql, uniqueId, v.Max-setting.SingleFile.MaxVersion+1)
	if err != nil {
		logger.Zap().Error("clean history version error", zap.Error(err), zap.String("sql", sql))
		return err
	}

	n, _ := res.RowsAffected()
	log.Printf("clean history version, unique_id:%s, max:%d, min:%d, affected:%d\n", uniqueId, v.Max, v.Min, n)

	// 扫描文件, 删除文件
	for i := v.Min; i > v.Min-setting.SingleFile.MaxVersion*2; i-- {
		full := fmt.Sprintf("%s/%s/%s/%s.%d.html", setting.SingleFile.Path, uniqueId[0:2], uniqueId[2:4], uniqueId, i)
		lite := fmt.Sprintf("%s/%s/%s/%s.%d.lite.html", setting.SingleFile.Path, uniqueId[0:2], uniqueId[2:4], uniqueId, i)
		_ = os.Remove(full)
		_ = os.Remove(lite)
		log.Printf("delete file, unique_id:%s, version:%d\n", uniqueId, i)
	}

	// todo 删除索引

	return nil
}
