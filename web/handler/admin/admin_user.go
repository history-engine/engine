package admin

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/user"
	"history-engine/engine/utils"
)

func UserList(c echo.Context) error {
	req := &model.UserListReq{}
	err := c.Bind(req)
	if err != nil {
		return utils.ApiError(c, model.ErrorParamParse)
	}

	list, code := user.List(c.Request().Context(), req)
	if code != model.Ok {
		return utils.ApiError(c, code)
	}

	return utils.ApiSuccess(c, list)
}
