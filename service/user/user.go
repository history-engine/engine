package user

import (
	"context"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/user"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/utils"
)

func Info(ctx context.Context, uid int) *ent.User {
	x := db.GetEngine()

	user, err := x.User.Get(ctx, uid)
	if err != nil {
		logger.Zap().Error("get user info err", zap.Error(err), zap.Int("uid", uid))
		return nil
	}

	return user
}

func Register(ctx context.Context, req *model.UserRegisterReq) (*ent.User, model.MsgCode) {
	x := db.GetEngine()

	count, err := x.User.Query().
		Where(
			user.Or(
				user.Username(req.Username),
				user.Email(req.Email),
			),
		).
		Count(ctx)
	if err != nil {
		logger.Zap().Error("check user exist err", zap.Error(err), zap.Any("req", req))
		return nil, model.Unknown
	}

	if count > 0 {
		return nil, model.ErrorUserExist
	}

	user, err := x.User.Create().
		SetUsername(req.Username).
		SetEmail(req.Email).
		SetPassword(utils.Md5str(req.Password)).
		Save(ctx)
	if err != nil {
		logger.Zap().Error("register err", zap.Error(err), zap.Any("req", req))
		return nil, model.Unknown
	}

	return user, model.Ok
}

func List(ctx context.Context, req *model.UserListReq) ([]*ent.User, model.MsgCode) {
	// todo
	return nil, model.Ok
}

func Create(ctx context.Context, req *model.UserCreateReq) model.MsgCode {
	// todo
	return model.Ok
}
