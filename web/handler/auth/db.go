package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/auth"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

// Password 密码登录
// TODO 后起期望提供 oauth等其他登录方式
func Password(c echo.Context) error {
	req := &model.PasswordLoginReq{}
	err := c.Bind(req)
	if err != nil {
		return utils.ApiError(c, model.ErrorParamParse)
	}

	u, err := auth.PasswordLogin(c.Request().Context(), req)
	if err != nil {
		panic(err)
	}

	if u == nil {
		return utils.ApiError(c, model.ErrorLoginFailed)
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
