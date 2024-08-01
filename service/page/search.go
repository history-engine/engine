package page

import (
	"context"
	"fmt"
	"history-engine/engine/ent"
	"history-engine/engine/model"
	"history-engine/engine/service/search"
	"history-engine/engine/setting"
)

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
			search.Engine().DelDocument(context.Background(), userId, item.DocId)
			continue
		}

		resp.Pages[k].Avatar = "https://avatars.akamai.steamstatic.com/6a9ae9c069cd4fff8bf954938727730cdb0fe27b.jpg"
		resp.Pages[k].Url = page.URL
		resp.Pages[k].Title = page.Title
		resp.Pages[k].Excerpt = page.Excerpt
		resp.Pages[k].Content = page.Content
		resp.Pages[k].Size = page.Size
		resp.Pages[k].Preview = setting.Web.Domain + "/page/view" + fmt.Sprintf("/%s.%d.html", page.UniqueID, page.Version)
		resp.Pages[k].Version = page.Version
		resp.Pages[k].CreatedAt = page.CreatedAt
	}

	return resp.Total, resp.Pages, err
}
