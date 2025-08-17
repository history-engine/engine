package host

import (
	"context"
	"fmt"
	"history-engine/engine/ent"
	entHost "history-engine/engine/ent/host"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
	"history-engine/engine/utils"
)

func Add(ctx context.Context, userId int64, host []string, Type int) error {
	defer func() {
		cache := localcache.GetEngine()
		key := fmt.Sprintf("host:all:%d:%d", userId, Type)
		cache.Delete(key)
	}()

	x := db.GetEngine()

	host = utils.FilterDuplicateDomains(host)
	create := make([]*ent.HostCreate, 0)
	for _, item := range host {
		exist, _ := x.Host.Query().Where(entHost.UserID(userId), entHost.Host(item), entHost.Type(Type)).Exist(ctx)
		if exist {
			continue
		}

		create = append(create, x.Host.Create().SetUserID(userId).SetType(Type).SetHost(item))
	}
	_, err := x.Host.CreateBulk(create...).Save(ctx)

	return err
}
