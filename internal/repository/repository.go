package repository

import (
	"context"

	"github.com/UraharaKiska/go-auth/internal/model"
	modelRepo "github.com/UraharaKiska/go-auth/internal/repository/user/model"
)

type UserRepository interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *modelRepo.UpdateUserInfo, id int64) error
	Delete(ctx context.Context, id int64) error
}
