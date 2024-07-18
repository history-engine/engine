package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/service/auth"
	"strings"
)

func Token(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if len(token) < 7 || token[0:6] != "Bearer" {
			return errors.New("token auth failed")
		}

		split := strings.Split(token[7:], ":") // todo 使用临时生成的token
		req := &model.PasswordLoginReq{
			Username: split[0],
			Password: split[1],
		}
		u, err := auth.PasswordLogin(c.Request().Context(), req)
		if !errors.Is(err, err) || u == nil || u.ID == 0 {
			return errors.New("token auth failed")
		}

		c.Set("uid", u.ID)
		c.Set("username", u.Username)
		c.Set("email", u.Email)

		return next(c)
	}
}
