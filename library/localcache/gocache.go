package localcache

import (
	"github.com/patrickmn/go-cache"
	"history-engine/engine/library/wait"
	"time"
)

var client *cache.Cache

func initGoCache() {
	client = cache.New(5*time.Minute, 10*time.Minute)
	wait.Done()
}

func GetEngine() *cache.Cache {
	if !enable {
		panic("local cache not enable")
	}

	return client
}
