package service

import (
	"context"
	"database/sql"
	"errors"
	"history-engine/engine/model"
	"history-engine/engine/utils"

	"github.com/jmoiron/sqlx"
)

// Auth auth service interface
type Auth interface {
	RegisterUser(ctx context.Context, username, email, passwd string) (*model.User, model.MsgCode)
	LoginWithPassword(ctx context.Context, username, password string) (*model.User, error)
}

type AuthImpl struct {
	db *sqlx.DB
}

func (a *AuthImpl) LoginWithPassword(ctx context.Context, username, passwd string) (*model.User, error) {
	user := &model.User{}
	// TODO 没有对输入name做校验，存在sql注入风险，优先可能使用orm封装的方法
	err := a.db.GetContext(ctx, user, "select * from user where username=? or email=? limit 1", username, username)
	if err == sql.ErrNoRows {
		// todo 用户不存在
		return nil, err
	} else if err != nil {
		panic(err)
	}
	if user.Id == 0 {
		return nil, errors.New("user not found")
	} else if user.Password != utils.Md5str(passwd) {
		return nil, errors.New("password is  wrong")
	}
	return user, nil
}

func (a *AuthImpl) RegisterUser(ctx context.Context, username, email, passwd string) (*model.User, model.MsgCode) {
	user := &model.User{}
	// TODO 需要对model层的一些方法进行封装，比如CRUD操作，裸写sql不够安全也不优雅
	err := a.db.GetContext(ctx, user, "select * from user where username=? or email=? limit 1", username, username)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if user.Id != 0 {
		// todo 用户已存在
		return nil, model.UserMailExist
	}
	user.Username = username
	user.Email = email
	user.Password = utils.Md5str(passwd)
	sql := "insert into user set " +
		"username=:username, email=:email, password=:password"
	res, err := a.db.NamedExecContext(ctx, sql, user)
	if err != nil {
		panic(err)
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return user, model.Ok
}
