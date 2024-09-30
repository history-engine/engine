package setting

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/ent"
	"history-engine/engine/model"
	"history-engine/engine/service/host"
	"history-engine/engine/utils"
)

func GetHost(ctx echo.Context) error {
	var userId = ctx.Get("uid").(int64)
	var page, limit int = 0, 0
	var keyword string = ""
	err := echo.QueryParamsBinder(ctx).
		Int("page", &page).
		Int("limit", &limit).
		String("keyword", &keyword).
		BindError()
	if err != nil {
		panic(err)
	}

	page = utils.Ternary(page <= 0, 1, page)
	limit = utils.Ternary(limit <= 0, 20, limit)

	total, hosts, err := host.Page(ctx.Request().Context(), userId, page, limit, keyword)
	resp := model.PageResponse[[]*ent.Host]{Total: total, Data: hosts}

	return utils.ApiSuccess(ctx, resp)
}

func SaveHost(ctx echo.Context) error {
	return nil
}