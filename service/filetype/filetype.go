package filetype

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/filetype"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
	"history-engine/engine/library/logger"
	"net/url"
	"strings"
	"time"
)

func Include(userId int64, filename string) bool {
	return suffixMatch(userId, filename, 1)
}

func Exclude(userId int64, filename string) bool {
	return suffixMatch(userId, filename, 2)
}

func suffixMatch(userId int64, filename string, Type int) bool {
	if filename == "" {
		return false
	}

	parse, err := url.Parse(filename)
	if err != nil {
		logger.Zap().Warn("parse filename err", zap.Error(err))
		return false
	}

	split := strings.Split(parse.Path, ".")
	if len(split) < 2 {
		return false
	}
	suffix := split[len(split)-1]

	filetypes, _ := All(context.Background(), userId, Type)
	if _, ok := filetypes[suffix]; ok {
		return true
	}

	return false
}

func All(ctx context.Context, userId int64, Type int) (map[string]struct{}, error) {
	cache := localcache.GetEngine()
	key := fmt.Sprintf("filetype:all:%d:%d", userId, Type)
	if cache, ok := cache.Get(key); ok {
		if list, ok := cache.(map[string]struct{}); ok {
			return list, nil
		}
	}

	x := db.GetEngine()
	list, err := x.FileType.Query().Select(filetype.FieldSuffix).Where(filetype.Type(Type)).All(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, err
	}

	filetypes := make(map[string]struct{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			filetypes[v.Suffix] = struct{}{}
		}
		cache.Set(key, filetypes, time.Hour*1)
	}

	return filetypes, nil
}
