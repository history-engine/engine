package logger

import (
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"

	"go.uber.org/zap"
)

var _zap *zap.Logger

func initZap() error {
	var err error
	var config zap.Config

	if setting.Log.Level == "debug" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	if setting.Log.File != "" {
		config.OutputPaths = []string{setting.Log.File}
		config.ErrorOutputPaths = []string{setting.Log.File}
	} else {
		config.OutputPaths = []string{"stdout"}
		config.ErrorOutputPaths = []string{"stderr"}
	}

	config.Encoding = "json"

	_zap, err = config.Build()
	_zap.Info("zap logger init success")

	wait.Done()
	return err
}

func Zap() *zap.Logger {
	if !enable {
		panic("logger not enable")
	}

	return _zap
}
