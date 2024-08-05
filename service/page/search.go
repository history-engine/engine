package page

import (
	"context"
	"fmt"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/service/icon"
	"history-engine/engine/service/search"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

func LatestList(ctx context.Context, userId int64, request model.SearchRequest) (int, []model.SearchResultPage, error) {
	x := db.GetEngine()

	total, err := x.Page.Query().Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	source, err := x.Page.
		Query().
		Where(page.UserID(userId)).
		Order(ent.Desc(page.FieldID)).
		Offset((request.Page - 1) * request.Limit).
		Limit(request.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	pages := make([]model.SearchResultPage, 0)
	for _, item := range source {
		row := model.SearchResultPage{
			Avatar:    icon.PublicUrl(ctx, item),
			Url:       item.URL,
			Title:     utils.Ternary(item.Title != "", item.Title, "无标题"),
			Excerpt:   item.Excerpt,
			Content:   item.Content,
			Size:      item.Size,
			Preview:   setting.Web.Domain + "/page/view" + fmt.Sprintf("/%s.%d.html", item.UniqueID, item.Version),
			DocId:     fmt.Sprintf("%s%d", item.UniqueID, item.Version),
			UniqueId:  item.UniqueID,
			Version:   item.Version,
			CreatedAt: item.CreatedAt,
		}
		pages = append(pages, row)
	}

	return total, pages, err
}

func Search(ctx context.Context, userId int64, request model.SearchRequest) (int, []model.SearchResultPage, error) {
	resp, err := search.Engine().Search(ctx, userId, request)
	if err != nil {
		return 0, nil, err
	}

	// 提取页面id，忽略版本
	docIdList := make([]string, 0)
	for _, item := range resp.Pages {
		docIdList = append(docIdList, item.UniqueId)
	}

	// 从数据获取页面信息，会额外回去没有搜索到的版本
	pages, err := BatchGetPage(ctx, docIdList)
	docMap := make(map[string]*ent.Page, 0)
	for _, v := range pages {
		docMap[fmt.Sprintf("%s%d", v.UniqueID, v.Version)] = v
	}

	// 保持搜索引擎的顺序，并从数据库补全剩余信息
	for k, item := range resp.Pages {
		page, ok := docMap[item.DocId]
		if !ok { // 搜索引擎存储，数据库不存在
			go search.Engine().DelDocument(context.Background(), userId, item.DocId)
			continue
		}

		resp.Pages[k].Avatar = icon.PublicUrl(ctx, page)
		resp.Pages[k].Url = page.URL
		resp.Pages[k].Title = utils.Ternary(page.Title != "", page.Title, "无标题")
		resp.Pages[k].Excerpt = page.Excerpt
		resp.Pages[k].Content = page.Content
		resp.Pages[k].Size = page.Size
		resp.Pages[k].Preview = setting.Web.Domain + "/page/view" + fmt.Sprintf("/%s.%d.html", page.UniqueID, page.Version)
		resp.Pages[k].Version = page.Version
		resp.Pages[k].CreatedAt = page.CreatedAt
	}

	return resp.Total, resp.Pages, err
}
