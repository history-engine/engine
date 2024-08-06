package page

import (
	"context"
	"errors"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/service/host"
	"net/url"
)

func Exclude(ctx context.Context, ident model.PageIdent) error {
	x := db.GetEngine()

	var err error
	var row *ent.Page
	if ident.Id > 0 {
		row, err = x.Page.Query().Where(page.ID(ident.Id)).First(ctx)
	} else {
		row, err = x.Page.Query().Where(page.UniqueID(ident.UniqueId), page.Version(ident.Version)).First(ctx)
	}

	if row == nil {
		return errors.New("page not found")
	}

	if err != nil {
		return nil
	}

	parsed, err := url.Parse(row.URL)
	if err != nil {
		return err
	}

	return host.Add(ctx, row.UserID, parsed.Host, 2)
}
