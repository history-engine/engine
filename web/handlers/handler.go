package handlers

import "history-engine/engine/service"

// Handler handler统一入口
type Handler interface {
	//TODO  SingleFile依赖暂时比较多，后面继续做
	// SingleFileHandler SingleFileHandler
	UserHandler
}

// HandlerImpl handler实现
type HandlerImpl struct {
	UserHandlerImpl
}

// NewHandler handler初始化
func NewHandlerImpl(svr service.ServiceInterface) *HandlerImpl {
	return &HandlerImpl{
		UserHandlerImpl: UserHandlerImpl{s: svr},
	}
}
