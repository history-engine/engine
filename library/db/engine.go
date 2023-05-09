package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"history-engine/engine/setting"
)

var x *sqlx.DB // todo 或许可以public直接调用

func InitEngine(ctx context.Context) error {
	db, err := sqlx.ConnectContext(ctx, setting.Database.Drive, setting.GetDSN())
	if err != nil {
		panic(err)
	}

	// todo 自定义mapper, 可以不用写db tag
	x = db.Unsafe()

	return nil
}

func GetEngine() *sqlx.DB {
	return x
}
