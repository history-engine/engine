package localcache

import (
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"time"
)

var client *cache.Cache

func initGoCache() {
	client = cache.New(5*time.Minute, 10*time.Minute)

	path := setting.Common.DataPath + "/local-cache.db"
	if utils.FileExist(path) {
		if err := client.LoadFile(path); err != nil {
			logger.Zap().Info("load local-cache file error:", zap.Error(err))
		}
	}

	wait.Done()
}

func GetEngine() *cache.Cache {
	if !enable {
		panic("local cache not enable")
	}

	return client
}

func persistent() {
	path := setting.Common.DataPath + "/local-cache.db"
	if err := client.SaveFile(path); err != nil {
		logger.Zap().Warn("save local-cache file error:", zap.Error(err))
	}
}
