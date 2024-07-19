package user

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/user"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"net/http"
	"time"
)

func Info(c echo.Context) error {
	if uid, ok := c.Get("uid").(int); ok {
		u := user.Info(c.Request().Context(), uid)
		return utils.ApiSuccess(c, u)
	}

	return utils.ApiError(c, model.ErrorLoginFailed)
}

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     setting.JwtKey,
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Second),
		Path:     "/",
		HttpOnly: true,
	})

	return utils.ApiSuccess(c, nil)
}

// Register 注册
func Register(c echo.Context) error {
	if !setting.Common.EnableRegister {
		return utils.ApiError(c, model.ErrorRegisterDisabled)
	}

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
		"uid":      u.ID,
		"username": u.Username,
		"email":    u.Email,
	})
	tokenString, err := token.SignedString(setting.JwtSecret)
	if err != nil {
		panic(err)
	}

	return utils.ApiSuccess(c, map[string]any{"jwt_token": tokenString})
}

// PasswordLogin 密码登录
func PasswordLogin(c echo.Context) error {
	req := &model.PasswordLoginReq{}
	err := c.Bind(req)
	if err != nil {
		return utils.ApiError(c, model.ErrorParamParse)
	}

	u, err := user.PasswordLogin(c.Request().Context(), req)
	if err != nil {
		return utils.ApiError(c, model.ErrorLoginFailed)
	}

	if u == nil {
		return utils.ApiError(c, model.ErrorLoginFailed)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      u.ID,
		"username": u.Username,
		"email":    u.Email,
	})
	tokenString, err := token.SignedString(setting.JwtSecret)
	if err != nil {
		panic(err)
	}

	c.SetCookie(&http.Cookie{
		Name:     setting.JwtKey,
		Value:    tokenString,
		Expires:  time.Now().Add(86400 * 24 * time.Second),
		Path:     "/",
		HttpOnly: true,
	})

	return utils.ApiSuccess(c, map[string]any{"jwt_token": tokenString, "user": u})
}
