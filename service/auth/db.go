package auth

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/user"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/utils"
)

func PasswordLogin(ctx context.Context, req *model.PasswordLoginReq) (*ent.User, error) {
	x := db.GetEngine()

	user, err := x.User.Query().
		Where(
			user.Or(
				user.Username(req.Username),
				user.Email(req.Username),
			),
			user.Password(utils.Md5str(req.Password)),
		).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, errors.New("user not found")
	} else if err != nil {
		logger.Zap().Error("password login err", zap.Error(err))
		return nil, err
	}

	return user, nil
}
