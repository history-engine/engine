package setting

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	pageService "history-engine/engine/service/page"
	"history-engine/engine/utils"
)

func GetPage(ctx echo.Context) error {
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

	req := model.SearchRequest{
		Query: keyword,
		Page:  page,
		Limit: limit,
	}

	total, pages, err := pageService.LatestList(ctx.Request().Context(), userId, req)
	resp := model.PageResponse[[]model.SearchResultPage]{Total: total, Data: pages}

	return utils.ApiSuccess(ctx, resp)
}

func DeletePage(ctx echo.Context) error {
	var userId = ctx.Get("uid").(int64)
	var id int64 = 0
	err := echo.QueryParamsBinder(ctx).Int64("id", &id).BindError()
	if err != nil {
		panic(err)
	}

	pageService.DeleteByIdent(ctx.Request().Context(), model.PageIdent{Id: id, UserId: userId})

	return utils.ApiSuccess(ctx, nil)
}
