package middleware

import (
	"history-engine/engine/model"
	"history-engine/engine/service/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		req := &model.PasswordLoginReq{
			Username: username,
			Password: password,
		}
		u, err := auth.PasswordLogin(c.Request().Context(), req)
		if u != nil && u.Id > 0 {
			c.Set("uid", u.Id)
			c.Set("username", u.Username)
			c.Set("email", u.Email)
			return true, nil
		}
		c.Redirect(302, "/login")
		return false, err
	})
}
