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

func AddHost(ctx echo.Context) error {
	var userId = ctx.Get("uid").(int64)

	row := ent.Host{}
	err := ctx.Bind(&row)
	if err != nil {
		panic(err)
	}

	err = host.Add(ctx.Request().Context(), userId, []string{row.Host}, row.Type)
	if err != nil {
		panic(err)
	}

	return utils.ApiSuccess(ctx, nil)
}

func SaveHost(ctx echo.Context) error {
	var userId = ctx.Get("uid").(int64)
	row := ent.Host{}
	err := ctx.Bind(&row)
	if err != nil {
		panic(err)
	}

	err = host.Edit(ctx.Request().Context(), userId, row)
	if err != nil {
		panic(err)
	}

	return utils.ApiSuccess(ctx, nil)
}

func DeleteHost(ctx echo.Context) error {
	var userId = ctx.Get("uid").(int64)
	var id int64 = 0
	err := echo.QueryParamsBinder(ctx).Int64("id", &id).BindError()
	if err != nil {
		panic(err)
	}

	host.Delete(ctx.Request().Context(), userId, id)

	return utils.ApiSuccess(ctx, nil)
}
