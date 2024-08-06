package page

import (
	"context"
	"errors"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/service/readability"
	"history-engine/engine/setting"
	"time"
)

func ParserPageWithId(ctx context.Context, id int64) error {
	return parserPage(ctx, model.PageIdent{Id: id})
}

func ParserPage(userId int64, uniqueId string, version int) error {
	row := model.PageIdent{
		UserId:   userId,
		UniqueId: uniqueId,
		Version:  version,
	}
	return parserPage(context.Background(), row)
}

// ParserPage 调用readability分析HTML文件、保存数据库
func parserPage(ctx context.Context, row model.PageIdent) error {
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

	fullPath := setting.SingleFile.HtmlPath + item.Path
	article, err := readability.Parser().Parse(fullPath)
	if err != nil && errors.Is(err, readability.ErrContentEmpty) {
		_, err2 := x.Page.Update().SetParsedAt(time.Now()).Where(page.ID(item.ID)).Save(ctx)
		return errors.Join(err, err2)
	} else if err != nil {
		return err
	}

	_, err = x.Page.Update().
		SetTitle(article.Title).
		SetExcerpt(article.Excerpt).
		SetContent(article.TextContent).
		SetParsedAt(time.Now()).
		Where(page.ID(item.ID)).
		Save(ctx)
	return err
}
