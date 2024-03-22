package db

import (
	"context"
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var x *sqlx.DB // todo 或许可以public直接调用

func initEngine() error {
	db, err := sqlx.ConnectContext(context.TODO(), setting.Database.Drive, setting.GetDSN())
	if err != nil {
		panic(err)
	}

	// todo 自定义mapper, 可以不用写db tag
	x = db.Unsafe()

	wait.Done()
	return nil
}

func GetEngine() *sqlx.DB {
	if !enable {
		panic("db not enable")
	}
	return x
}
