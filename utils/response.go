package utils

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/model"
	"net/http"
)

func ApiResponse(ctx echo.Context, code int, msg string, data any) error {
	return ctx.JSON(http.StatusOK, model.ApiResponse{
		Code:    0,
		Message: msg,
		Data:    data,
	})
}

func ApiSuccess(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, model.ApiResponse{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

func ApiError(ctx echo.Context, code model.MsgCode) error {
	var ok bool
	msg := "ok"
	if msg, ok = model.MsgTable[code]; !ok {
		msg = model.MsgTable[model.Unknown]
	}

	return ctx.JSON(http.StatusOK, model.ApiResponse{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func ApiServerError(ctx echo.Context, code model.MsgCode) error {
	var ok bool
	msg := "ok"
	if msg, ok = model.MsgTable[code]; !ok {
		msg = model.MsgTable[model.Unknown]
	}

	return ctx.JSON(http.StatusInternalServerError, model.ApiResponse{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func ApiErrorWithData(ctx echo.Context, code model.MsgCode, data any) error {
	var ok bool
	msg := "ok"
	if msg, ok = model.MsgTable[code]; !ok {
		msg = model.MsgTable[model.Unknown]
	}

	return ctx.JSON(200, model.ApiResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
