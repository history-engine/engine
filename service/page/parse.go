package page

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/service/readability"
	"history-engine/engine/setting"
	"os"
	"time"
)

func ParserPageWithId(id int64) error {
	return parserPage(context.Background(), model.PageParse{Id: id})
}

func ParserPage(userId int64, uniqueId string, version int) error {
	row := model.PageParse{
		UserId:   userId,
		UniqueId: uniqueId,
		Version:  version,
	}
	return parserPage(context.Background(), row)
}

// ParserPage 调用readability分析HTML文件、保存数据库
func parserPage(ctx context.Context, row model.PageParse) error {
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

	if item.Content != "" {
		return nil
	}

	fullPath := setting.SingleFile.HtmlPath + item.Path
	if f, err := os.Stat(fullPath); err != nil {
		return err
	} else if f.Size() > int64(setting.SingleFile.MaxSize) {
		return errors.New(fmt.Sprintf("File size exceeds threshold: %d", setting.SingleFile.MaxSize))
	}

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
