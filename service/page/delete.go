package page

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/service/search"
	"history-engine/engine/setting"
	"os"
)

func DeleteByIdent(ctx context.Context, ident model.PageIdent) error {
	var err error
	var row *ent.Page
	x := db.GetEngine()
	if ident.Id > 0 && ident.UserId > 0 {
		row, err = x.Page.Query().Where(page.ID(ident.Id), page.UserID(ident.UserId)).First(ctx)
	} else if ident.UserId > 0 && ident.UniqueId != "" && ident.Version > 0 {
		row, err = x.Page.Query().
			Where(
				page.UserID(ident.UserId),
				page.UniqueID(ident.UniqueId),
				page.Version(ident.Version),
			).
			First(ctx)
	} else {
		return errors.New("ident invalid")
	}

	if err != nil {
		return err
	}

	return Delete(ctx, row)
}

func Delete(ctx context.Context, row *ent.Page) error {
	if row == nil {
		return errors.New("page empty")
	}

	if err := os.Remove(setting.Common.HtmlPath + row.Path); err != nil && !errors.As(err, &os.ErrNotExist) {
		return err
	}

	docId := fmt.Sprintf("%s%d", row.UniqueID, row.Version)
	if err := search.Engine().DelDocument(ctx, row.UserID, docId); err != nil {
		return err
	}

	return db.GetEngine().Page.DeleteOneID(row.ID).Exec(ctx)
}
