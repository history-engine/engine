package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/setting"
	"history-engine/engine/web/handler/singlefile"
	"history-engine/engine/web/middleware"
)

// 注册页面相关路由
func singleFileRouteRegister(r *echo.Group) {
	r.Add("POST", "/rest", singlefile.RestSave, middleware.Token)

	r.Use(middleware.BasicAuth())
	e := singlefile.NewEndpoint("/", setting.SingleFile.Path)
	r.Add("PUT", "/html/:file", e.Put)
	r.Add("OPTIONS", "/html/:file", e.Cover)
	r.Add("MKCOL", "/html/:file", e.Cover)
	r.Add("PROPFIND", "/html/:file", e.Cover)
	r.Add("HEAD", "/html/:file", e.Head)
}

/* func basicAuth(username, password string, c echo.Context) (bool, error) {
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

	return false, err
} */
