package redis

import (
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"

	"github.com/redis/go-redis/v9"
)

var _redis *redis.Client

func initEngine() error {
	_redis = redis.NewClient(&redis.Options{
		Addr:       setting.Redis.Addr,
		Password:   setting.Redis.Password,
		PoolFIFO:   true,
		ClientName: "history-engine",
		DB:         setting.Redis.DB,
		// PoolSize:       10,
		// MaxRetries:     3,
		// MaxIdleConns:   0,
		// MaxActiveConns: 0,
	})
	wait.Done()
	return nil
}

func GetEngine() *redis.Client {
	if !enable {
		panic("redis not enable")
	}
	return _redis
}
