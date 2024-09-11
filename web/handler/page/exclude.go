package page

import (
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/utils"

	"github.com/labstack/echo/v4"
)

func Exclude(c echo.Context) error {
	req := model.ExcludeRequest{UserId: c.Get("uid").(int64)}
	if err := c.Bind(&req); err != nil {
		return utils.ApiResponse(c, 500, err.Error(), req)
	}

	if req.UserId == 0 || req.UniqueId == "" || req.Version == 0 || len(req.Domains) == 0 {
		return utils.ApiResponse(c, 500, "args empty", req)
	}

	if err := page.Exclude(c.Request().Context(), req); err != nil {
		return utils.ApiResponse(c, 500, err.Error(), nil)
	}

	return utils.ApiResponse(c, 200, "ok", nil)
}
