package page

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	servicePage "history-engine/engine/service/page"
	"history-engine/engine/utils"
)

func Versions(c echo.Context) error {
	var page, limit, uniqueId = 0, 0, ""

	userId := c.Get("uid").(int64)
	err := echo.QueryParamsBinder(c).
		Int("page", &page).
		Int("limit", &limit).
		String("unique_id", &uniqueId).
		FailFast(true).
		BindError()
	if err != nil {
		return c.JSON(400, model.ApiResponse{Code: -1, Message: err.Error()})
	}

	if page < 1 {
		page = 1
	}

	if limit == 0 {
		limit = 20
	}

	resp := model.SearchResponse{}
	resp.Total, resp.Pages, err = servicePage.Versions(c.Request().Context(), userId, uniqueId, page, limit)

	if err != nil {
		return c.JSON(200, model.ApiResponse{Code: -1, Message: err.Error()})
	}

	return utils.ApiSuccess(c, resp)
}
