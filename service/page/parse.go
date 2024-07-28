package page

import (
	"context"
	"fmt"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/service/readability"
	"history-engine/engine/service/zincsearch"
	"history-engine/engine/setting"
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

// ParserPage 调用readability分析HTML文件，添加到ZincSearch、保存数据库
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

	article := &readability.Article{}
	if item.Excerpt == "" && item.Content == "" {
		article, err = readability.Parser().Parse(setting.SingleFile.HtmlPath + item.Path)
		if err != nil {
			return err
		}

		_, err = x.Page.Update().
			SetTitle(article.Title).
			SetExcerpt(article.Excerpt).
			SetContent(article.TextContent).
			Where(page.ID(item.ID)).
			Save(ctx)
		if err != nil {
			return err
		}

		item.Title = article.Title
		item.Excerpt = article.Excerpt
		item.Content = article.TextContent
	}

	zincId := fmt.Sprintf("%s%d", item.UniqueID, item.Version)
	zincDoc := &model.ZincDocument{
		Url:     item.URL,
		Title:   item.Title,
		Excerpt: item.Excerpt,
		Content: item.Content,
	}
	if err = zincsearch.PutDocument(item.UserID, zincId, zincDoc); err != nil {
		return err
	}

	_, err = x.Page.Update().SetIndexedAt(time.Now()).Where(page.ID(item.ID)).Save(ctx)

	return err
}
