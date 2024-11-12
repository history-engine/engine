package setting

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/ent"
	serviceSetting "history-engine/engine/service/setting"
	"history-engine/engine/utils"
)

func GetStorage(ctx echo.Context) error {
	var userId = ctx.Get("uid").(int64)
	data, err := serviceSetting.GetSetting(ctx.Request().Context(), userId)
	if err != nil {
		panic(err)
	}

	return utils.ApiSuccess(ctx, data)
}

func SaveStorage(ctx echo.Context) error {
	var userId = ctx.Get("uid").(int64)
	row := &ent.Setting{}
	err := ctx.Bind(&row)
	if err != nil {
		panic(err)
	}

	err = serviceSetting.Save(ctx.Request().Context(), userId, row)
	if err != nil {
		panic(err)
	}

	return utils.ApiSuccess(ctx, nil)

	return utils.ApiSuccess(ctx, nil)
}
