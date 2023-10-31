package logger

import (
	"context"
	"go.uber.org/zap"
	"history-engine/engine/setting"
)

var _zap *zap.Logger

func InitZap(ctx context.Context) error {
	var err error
	var config zap.Config

	if setting.Log.Level == "debug" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	if setting.Log.Path != "" {
		config.OutputPaths = []string{setting.Log.Path}
		config.ErrorOutputPaths = []string{setting.Log.Path}
	} else {
		config.OutputPaths = []string{"stdout"}
		config.ErrorOutputPaths = []string{"stderr"}
	}

	config.Encoding = "json"

	_zap, err = config.Build()
	_zap.Info("zap logger init success")

	return err
}

func Zap() *zap.Logger {
	if _zap == nil {
		if err := InitZap(context.TODO()); err != nil {
			panic(err)
		}
	}
	return _zap
}
