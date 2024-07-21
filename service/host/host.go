package host

import (
	"context"
	"fmt"
	"github.com/tidwall/match"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/host"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
	"history-engine/engine/library/logger"
	"net/url"
	"regexp"
	"time"
)

func Include(userId int64, address string) bool {
	return hostMatch(userId, address, 1)
}

func Exclude(userId int64, address string) bool {
	return hostMatch(userId, address, 2)
}

func hostMatch(userId int64, address string, Type int) bool {
	if address == "" {
		return false
	}

	parse, err := url.Parse(address)
	if err != nil {
		logger.Zap().Warn("parse host err", zap.Error(err))
		return false
	}

	var find = false
	hosts, _ := All(context.Background(), userId, Type)
	for _, item := range hosts {
		if len(item) > 7 && item[0:7] == "regexp:" {
			compile, err := regexp.Compile(item[7:])
			if err != nil {
				continue
			}
			find = compile.MatchString(parse.Host)
		} else {
			find = match.Match(parse.Host, item)
		}

		if find {
			break
		}
	}

	return find
}

func All(ctx context.Context, userId int64, Type int) ([]string, error) {
	cache := localcache.GetEngine()
	key := fmt.Sprintf("host:all:%d:%d", userId, Type)
	if cache, ok := cache.Get(key); ok {
		if list, ok := cache.([]string); ok {
			return list, nil
		}
	}

	x := db.GetEngine()
	list, err := x.Host.Query().Select(host.FieldHost).Where(host.Type(Type)).All(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, err
	}

	hosts := make([]string, 0)
	if len(list) > 0 {
		for _, v := range list {
			hosts = append(hosts, v.Host)
		}
		cache.Set(key, hosts, time.Hour*1)
	}

	return hosts, nil
}
