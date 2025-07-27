package service

import (
	"context"

	"github.com/UraharaKiska/go-auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *model.UserUpdateInfo, id int64) error
	Delete(ctx context.Context, id int64) error
}
