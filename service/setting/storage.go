package setting

import (
	"context"
	"history-engine/engine/ent"
	"history-engine/engine/ent/setting"
	"history-engine/engine/library/db"
	"history-engine/engine/utils"
)

func GetSetting(ctx context.Context, userId int64) (*ent.Setting, error) {
	s, err := db.GetEngine().Setting.Query().Where(setting.UserID(userId)).First(ctx)

	// 默认值
	if s != nil {
		s.MaxSize = utils.Ternary(s.MaxSize <= 0, 5, s.MaxSize)
		s.MinVersionInterval = utils.Ternary(s.MinVersionInterval <= 0, 86400, s.MinVersionInterval)
		s.MinSize = utils.Ternary(s.MinSize <= 0, 2048, s.MinSize)
		s.MaxSize = utils.Ternary(s.MaxSize <= 0, 20971520, s.MaxSize)
	}

	return s, err
}

func Save(ctx context.Context, userId int64, row *ent.Setting) error {
	x := db.GetEngine().Setting

	err := x.Update().
		Where(setting.UserID(userId)).
		SetMaxVersion(row.MaxVersion).
		SetMinVersionInterval(row.MinVersionInterval).
		SetMinSize(row.MinSize).
		SetMaxSize(row.MaxSize).
		Exec(ctx)

	return err
}
