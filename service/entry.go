package service

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type ServiceInterface interface {
	Auth
}

// Impl service对外接口
type Impl struct {
	AuthImpl
}

// NewServiceImpl new service impl
func NewServiceImpl(ctx context.Context, db *sqlx.DB) *Impl {
	return &Impl{
		AuthImpl: AuthImpl{db: db},
	}
}
