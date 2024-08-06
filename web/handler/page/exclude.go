package page

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/utils"
)

func Exclude(c echo.Context) error {
	ident := model.PageIdent{UserId: c.Get("uid").(int64)}
	if err := c.Bind(&ident); err != nil {
		return utils.ApiResponse(c, 500, err.Error(), ident)
	}

	if ident.Id == 0 && ident.UserId == 0 && ident.UniqueId == "" && ident.Version == 0 {
		return utils.ApiResponse(c, 500, "args empty", ident)
	}

	if err := page.Exclude(c.Request().Context(), ident); err != nil {
		return utils.ApiResponse(c, 500, err.Error(), nil)
	}

	return utils.ApiResponse(c, 200, "ok", nil)
}
