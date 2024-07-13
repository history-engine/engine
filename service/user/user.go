package user

import (
	"context"
	"database/sql"
	"errors"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/utils"
)

func Info(ctx context.Context, uid int64) *model.User {
	x := db.GetEngine()
	user := &model.User{}
	err := x.GetContext(ctx, user, "select * from user where id=?", uid)
	if err != nil {
		panic(err)
	}

	return user
}

func Register(ctx context.Context, req *model.UserRegisterReq) (*model.User, model.MsgCode) {
	x := db.GetEngine()
	user := &model.User{}
	err := x.GetContext(ctx, user, "select * from user where username=? or email=? limit 1", req.Username, req.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}

	if user.Id != 0 {
		// todo 用户已存在
		return nil, model.ErrorUserExist
	}

	user.Username = req.Username
	user.Email = req.Email
	user.Password = utils.Md5str(req.Password)
	query := "insert into user set " +
		"username=:username, email=:email, password=:password"
	res, err := x.NamedExecContext(ctx, query, user)
	if err != nil {
		panic(err)
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return user, model.Ok
}

func List(ctx context.Context, req *model.UserListReq) ([]model.User, model.MsgCode) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Rows <= 0 {
		req.Rows = 20
	}

	x := db.GetEngine()
	users := make([]model.User, 0)
	err := x.SelectContext(ctx, &users, "select * from user limit ?,?", (req.Page-1)*20, 20)
	if err != nil {
		panic(err)
	}

	return users, model.Ok
}

func Create(ctx context.Context, req *model.UserCreateReq) model.MsgCode {
	x := db.GetEngine()

	// 重复检查
	query := "select * from user where username=? or email=? limit 1"
	user := &model.User{}
	err := x.GetContext(ctx, user, query, req.Username, req.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}

	if user.Id != 0 {
		return model.ErrorUserExist
	}

	user = &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: utils.Md5str(req.Password),
	}
	query = "insert into user set " +
		"username=:username, email=:email, password=:password"
	_, err = x.NamedExecContext(ctx, query, user)
	if err != nil {
		panic(err)
	}

	return model.Ok
}
