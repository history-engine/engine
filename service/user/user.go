package user

import (
	"context"
	"database/sql"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/utils"
)

func Register(ctx context.Context, req *model.UserRegisterReq) (*model.User, model.MsgCode) {
	x := db.GetEngine()
	user := &model.User{}
	err := x.GetContext(ctx, user, "select * from user where username=? or email=? limit 1", req.Username, req.Username)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if user.Id != 0 {
		// todo 用户已存在
		return nil, model.UserMailExist
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
