package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

// JwtAuth Jwt验证
// todo 可能不需要上jwt这种验证方式
func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get jwt token from query string or header or cookie
		jwtToken := c.QueryParam(setting.JwtKey)
		if jwtToken == "" {
			jwtToken = c.Request().Header.Get(setting.JwtKey)
			if jwtToken == "" {
				cookie, err := c.Cookie(setting.JwtKey)
				if err == nil && cookie != nil {
					jwtToken = cookie.Value
				}
			}
		}
		if jwtToken == "" {
			return utils.ApiError(c, model.ErrorLoginFailed)
		}

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return setting.JwtSecret, nil
		})
		if err == nil || token == nil || !token.Valid {
			return utils.ApiError(c, model.ErrorLoginFailed)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("jwt-claims", claims)
			if uid, ok := claims["uid"]; ok {
				c.Set("uid", utils.Str2Int(uid.(string)))
			}
		}

		return next(c)
	}
}
