package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"
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
	}

	config.Encoding = setting.Log.Format
	config.DisableCaller = true
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncoderConfig.ConsoleSeparator = " "

	_zap, err = config.Build()

	wait.Done()
	return err
}

func Zap() *zap.Logger {
	if !enable {
		panic("logger not enable")
	}

	return _zap
}
