package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/library/logger"
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"
	"log"
	_ "modernc.org/sqlite"
)

var client *ent.Client

func initEngine() {
	var err error

	opt := ent.Log(func(a ...any) {
		logger.Zap().Info("ent debug", zap.Any("info", a))
	})

	client, err = ent.Open(setting.Database.Drive, setting.GetDSN(), opt)
	if err != nil {
		log.Fatalf("connect to %s err: %v\b", setting.Database.Drive, err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	wait.Done()
}

func GetEngine() *ent.Client {
	if !enable {
		panic("db not enable")
	}

	if setting.Common.Env == "dev" && setting.Database.PrintSql {
		return client.Debug()
	}

	return client
}
