package page

import (
	"context"
	"errors"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
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

	return host.Add(ctx, params.UserId, params.Domains, 2)
}
