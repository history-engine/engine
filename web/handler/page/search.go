package page

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/service/zincsearch"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"time"
)

func Search(c echo.Context) error {
	req := model.SearchPage{}
	if err := c.Bind(&req); err != nil {
		return c.String(400, err.Error())
	}

	if req.Limit == 0 {
		req.Limit = 20
	}

	if req.StartTime.IsZero() {
		req.StartTime = time.Now().AddDate(0, 0, -30)
	}

	if req.EndTime.IsZero() {
		req.EndTime = time.Now()
	}

	resp, err := zincsearch.EsSearch(req)
	if err != nil {
		return c.String(500, err.Error())
	}

	pages, err := page.BatchGetPage(c.Request().Context(), resp)
	if err != nil {
		panic(err)
	}

	versions := map[string][]model.Page{}
	for _, v := range pages {
		v.FullPath = setting.Web.Domain + "/page/preview" + v.FullPath
		if _, ok := versions[v.UniqueId]; !ok {
			versions[v.UniqueId] = []model.Page{}
		}
		versions[v.UniqueId] = append(versions[v.UniqueId], v)
	}

	return utils.ApiSuccess(c, versions)
}
