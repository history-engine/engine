package db

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"history-engine/engine/ent"
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"
	"log"
)

var client *ent.Client

func initEngine() {
	var err error

	client, err = ent.Open(setting.Database.Drive, setting.GetDSN())
	if err != nil {
		log.Fatalf("connect to %s err: %v\b", setting.Database.Drive, err)
	}

	wait.Done()
}

func GetEngine() *ent.Client {
	if !enable {
		panic("db not enable")
	}

	if setting.Common.Env == "dev" {
		return client.Debug()
	}

	return client
}
