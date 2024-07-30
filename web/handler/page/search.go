package page

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/utils"
	"time"
)

func Search(c echo.Context) error {
	userId := c.Get("uid").(int64)

	req := model.SearchRequest{}
	if err := c.Bind(&req); err != nil {
		return c.String(400, err.Error())
	}

	if req.Page < 1 {
		req.Page = 1
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

	result, err := page.Search(c.Request().Context(), userId, req)
	if err != nil {
		panic(err)
	}

	return utils.ApiSuccess(c, model.SearchResponse{Total: result.Total, Pages: result.Pages})
}
