package auth

import (
	"context"
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/utils"
)

func PasswordLogin(ctx context.Context, req *model.PasswordLoginReq) (*model.User, error) {
	x := db.GetEngine()

	query := "select * from user where username=? or email=? limit 1"
	user := &model.User{}
	err := x.GetContext(ctx, user, query, req.Username, req.Username)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		logger.Zap().Error("search user err", zap.Error(err), zap.String("sql", query))
		return nil, err
	}

	if user.Id == 0 || user.Password != utils.Md5str(req.Password) {
		return nil, errors.New("user not found")
	}

	return user, nil
}
