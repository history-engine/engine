package icon

import (
	"context"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/icon"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
	"history-engine/engine/library/logger"
	"history-engine/engine/setting"
	"net/url"
	"time"
)

func PublicUrl(ctx context.Context, page *ent.Page) string {
	var defaultIcon = setting.Web.Domain + "/page/icon/default.svg"

	parsed, err := url.Parse(page.URL)
	if err != nil {
		logger.Zap().Error("parse page url err", zap.Error(err), zap.String("url", page.URL))
		return defaultIcon
	}

	icons, err := All(ctx)
	if err != nil || len(icons) == 0 {
		return defaultIcon
	}

	if k, ok := icons[parsed.Host]; ok && k != "" {
		return setting.Web.Domain + "/page/icon/" + k
	}

	return defaultIcon
}

func All(ctx context.Context) (map[string]string, error) {
	cache := localcache.GetEngine()
	key := "icon:all"
	if cache, ok := cache.Get(key); ok {
		if list, ok := cache.(map[string]string); ok {
			return list, nil
		}
	}

	x := db.GetEngine()
	list, err := x.Icon.Query().Select(icon.FieldHost, icon.FieldPath).All(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, err
	}

	icons := make(map[string]string, 0)
	if len(list) > 0 {
		for _, v := range list {
			icons[v.Host] = v.Path
		}
		cache.Set(key, icons, time.Hour*1)
	}

	return icons, nil
}
