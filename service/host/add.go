package host

import (
	"context"
	"fmt"
	entHost "history-engine/engine/ent/host"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
)

func Add(ctx context.Context, userId int64, host string, Type int) error {
	defer func() {
		cache := localcache.GetEngine()
		key := fmt.Sprintf("host:all:%d:%d", userId, Type)
		cache.Delete(key)
	}()

	x := db.GetEngine()

	exist, _ := x.Host.Query().Where(entHost.UserID(userId), entHost.Host(host), entHost.Type(Type)).Exist(ctx)
	if exist {
		return nil
	}

	_, err := x.Host.Create().SetUserID(userId).SetType(Type).SetHost(host).Save(ctx)

	return err
}
