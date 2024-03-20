package handlers

import (
	"history-engine/engine/model"
	"history-engine/engine/service"
	"history-engine/engine/setting"
	"history-engine/engine/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// UserHandler user handler interface
type UserHandler interface {
	RegisterUser(c echo.Context) error
	Login(c echo.Context) error
}

type UserHandlerImpl struct {
	s service.ServiceInterface
}

// Register 注册
func (impl *UserHandlerImpl) RegisterUser(c echo.Context) error {
	req := &model.UserRegisterReq{}
	err := c.Bind(req)
	if err != nil {
		return utils.ApiError(c, model.ErrorParamParse)
	}

	u, code := impl.s.RegisterUser(c.Request().Context(), req.Username, req.Password, req.Email)
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

// Password 密码登录
// TODO 后起期望提供 oauth等其他登录方式
func (impl *UserHandlerImpl) Login(c echo.Context) error {
	req := &model.PasswordLoginReq{}
	err := c.Bind(req)
	if err != nil {
		return utils.ApiError(c, model.ErrorParamParse)
	}

	u, err := impl.s.LoginWithPassword(c.Request().Context(), req.Username, req.Password)
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
