package auth

import (
	"context"
	"database/sql"
	"errors"
	"history-engine/engine/library/db"
	"history-engine/engine/model"
	"history-engine/engine/utils"
)

func PasswordLogin(ctx context.Context, req *model.PasswordLoginReq) (*model.User, error) {
	x := db.GetEngine()

	user := &model.User{}
	err := x.GetContext(ctx, user, "select * from user where username=? or email=? limit 1", req.Username, req.Username)
	if err == sql.ErrNoRows {
		// todo 用户不存在
		return nil, err
	} else if err != nil {
		panic(err)
	}

	if user.Id == 0 || user.Password != utils.Md5str(req.Password) {
		return nil, errors.New("user not found")
	}

	return user, nil
}
