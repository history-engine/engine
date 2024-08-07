package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"history-engine/engine/model"
	"history-engine/engine/service/user"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		req := &model.PasswordLoginReq{
			Username: username,
			Password: password,
		}
		u, err := user.PasswordLogin(c.Request().Context(), req)
		if u != nil && u.ID > 0 {
			c.Set("uid", u.ID)
			c.Set("username", u.Username)
			c.Set("email", u.Email)
			return true, nil
		}

		return false, err
	})
}
