package page

import (
	"context"
	"fmt"
	"history-engine/engine/ent"
	"history-engine/engine/model"
	"history-engine/engine/service/icon"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

func EntPage2SearchResultPage(ctx context.Context, entPage *ent.Page) model.SearchResultPage {
	return model.SearchResultPage{
		Id:       entPage.ID,
		Avatar:   icon.PublicUrl(ctx, entPage),
		Url:      entPage.URL,
		Title:    utils.Ternary(entPage.Title != "", entPage.Title, "无标题"),
		Excerpt:  entPage.Excerpt,
		Content:  entPage.Content,
		Size:     utils.SizeFormat(entPage.Size),
		Preview:  setting.Web.Domain + "/page/view" + fmt.Sprintf("/%s.%d.html", entPage.UniqueID, entPage.Version),
		DocId:    fmt.Sprintf("%s%d", entPage.UniqueID, entPage.Version),
		UniqueId: entPage.UniqueID,
		Version:  entPage.Version,
		Time:     entPage.CreatedAt.Format("2006-01-02 15:05"),
	}
}
