package page

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/utils"
	"net/http"
)

// RestSave rest api 方式保存HTML
func RestSave(c echo.Context) error {
	err := c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	url := c.FormValue("url")
	html, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	src, err := html.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	hi := &model.HtmlInfo{
		Url:      url,
		Suffix:   utils.FileSuffix(url),
		Size:     int(html.Size),
		Sha1:     utils.Sha1Str(url),
		UserId:   c.Get("uid").(int64),
		IoReader: src,
	}

	if ok, err := page.Filter(hi); !ok {
		logger.Zap().Info(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := page.SavePage(c.Request().Context(), hi); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, nil)
}

// WebDavPreSave WebDAV方式保存HTML 前置检查
func WebDavPreSave(c echo.Context) error {
	hi := page.ParseHtmlInfo(c.Param("file"))
	hi.UserId = c.Get("uid").(int64)

	if ok, err := page.Filter(hi); !ok {
		logger.Zap().Info(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNotFound, nil)
}

// WebDavSave WebDAV方式保存HTML
func WebDavSave(c echo.Context) error {
	hi := page.ParseHtmlInfo(c.Param("file"))
	if hi == nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	hi.UserId = c.Get("uid").(int64)
	hi.Size = int(c.Request().ContentLength)
	hi.IoReader = c.Request().Body

	if err := page.SavePage(c.Request().Context(), hi); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, nil)
}
