package user

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/user"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

// Register 注册
func Register(c echo.Context) error {
	req := &model.UserRegisterReq{}
	err := c.Bind(req)
	if err != nil {
		return utils.ApiError(c, model.ErrorParamParse)
	}

	u, code := user.Register(c.Request().Context(), req)
	if code != model.Ok {
		return utils.ApiError(c, code)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      u.Id,
		"username": u.Username,
		"email":    u.Email,
	})
	tokenString, err := token.SignedString(setting.JwtSecret)
	if err != nil {
		panic(err)
	}

	return utils.ApiSuccess(c, map[string]any{"jwt_token": tokenString})
}
