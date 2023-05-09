package page

import (
	"context"
	"fmt"
	"history-engine/engine/library/db"
	"history-engine/engine/setting"
	"log"
	"os"
)

// NextVersion 获取下一个版本号
func NextVersion(ctx context.Context, uniqueId string) int {
	var version int
	x := db.GetEngine()
	err := x.GetContext(ctx, &version, "select max(version) from page where unique_id=?", uniqueId)
	if err != nil {
		return 1
	}

	return version + 1
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
	err := x.GetContext(ctx, &v, "select min(version) as min, max(version) as max from page where unique_id=?", uniqueId)
	if err != nil {
		// todo
		return err
	}

	if v.Max-v.Min < setting.SingleFile.MaxVersion {
		return nil
	}

	// todo 改成全部查出来再删除比较好因为还要删除文件和索引
	res, err := x.ExecContext(ctx, "delete from page where unique_id=? and version < ?", uniqueId, v.Max-setting.SingleFile.MaxVersion+1)
	if err != nil {
		// todo
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
