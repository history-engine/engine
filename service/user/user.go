package user

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/user"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/zincsearch"
	"history-engine/engine/utils"
)

func Info(ctx context.Context, uid int64) *ent.User {
	x := db.GetEngine()

	user, err := x.User.Get(ctx, uid)
	if err != nil {
		logger.Zap().Error("get user info err", zap.Error(err), zap.Int64("uid", uid))
		return nil
	}

	return user
}

// Register todo 不再返回model.MsgCode
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
		return nil, model.ErrorEmpty
	}

	if err := zincsearch.CreateIndex(user.ID); err != nil {
		return nil, model.ErrorInternal
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
