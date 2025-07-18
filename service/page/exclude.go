package page

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/host"
)

func Exclude(ctx context.Context, params model.ExcludeRequest) error {
	x := db.GetEngine()

	exist, err := x.Page.Query().
		Where(
			page.UserID(params.UserId),
			page.UniqueID(params.UniqueId),
			page.Version(params.Version),
		).
		Exist(ctx)
	if !exist {
		return errors.New("page not found")
	}

	if err != nil {
		return err
	}

	go func() {
		n, err := DeleteByHost(context.Background(), params.Domains...)
		if err != nil {
			logger.Zap().Error("delete ignore page err", zap.Error(err), zap.Any("domains", params.Domains))
		} else {
			logger.Zap().Info("delete ignore page", zap.Int("rows", n))
		}
	}()

	return host.Add(ctx, params.UserId, params.Domains, 2)
}
