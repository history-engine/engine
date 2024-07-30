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
	"time"
)

func PutIndexWithId(id int64) error {
	return putIndex(context.Background(), model.PageParse{Id: id})
}

func PutIndex(userId int64, uniqueId string, version int) error {
	row := model.PageParse{
		UserId:   userId,
		UniqueId: uniqueId,
		Version:  version,
	}
	return putIndex(context.Background(), row)
}

func putIndex(ctx context.Context, row model.PageParse) error {
	x := db.GetEngine()

	var err error
	item := &ent.Page{}
	if row.Id > 0 {
		item, err = x.Page.Get(ctx, row.Id)
	} else {
		item, err = x.Page.Query().
			Where(page.UserID(row.UserId), page.UniqueID(row.UniqueId), page.Version(row.Version)).
			First(ctx)
	}
	if err != nil {
		return err
	}

	if item.Title == "" && item.Excerpt == "" && item.Content == "" {
		return errors.New("page not parsed")
	}

	docId := fmt.Sprintf("%s%d", item.UniqueID, item.Version)
	doc := &model.SearchEngineDocument{
		Url:     item.URL,
		Title:   item.Title,
		Excerpt: item.Excerpt,
		Content: item.Content,
	}

	if err := search.Engine().PutDocument(ctx, item.UserID, docId, doc); err != nil {
		return err
	}

	_, err = x.Page.Update().SetIndexedAt(time.Now()).Where(page.ID(item.ID)).Save(ctx)
	return err
}
