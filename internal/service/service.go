package service

import (
	"context"

	"github.com/UraharaKiska/go-auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *model.UserUpdateInfo, id int64) error
	Delete(ctx context.Context, id int64) error
}

type AuthService interface {
	Login(ctx context.Context, user *model.UserInfo) (string, error)
	GetRefreshToken(ctx context.Context, token string) (string, error)
	GetAccessToken(ctx context.Context, token string) (string, error)
}

type AccessService interface {
	Check(ctx context.Context, endpoint string) (error)
}

