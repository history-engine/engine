package page

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/utils"
	"time"
)

func Search(c echo.Context) error {
	var err error
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

	resp := model.SearchResponse{}
	if req.Query == "" {
		resp.Total, resp.Pages, err = page.LatestList(c.Request().Context(), userId, req)
	} else {
		resp.Total, resp.Pages, err = page.Search(c.Request().Context(), userId, req)
	}
	if err != nil {
		return c.JSON(200, model.ApiResponse{Code: -1, Message: err.Error()})
	}

	return utils.ApiSuccess(c, resp)
}
