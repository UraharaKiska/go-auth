package repository

import (
	"context"

	"github.com/UraharaKiska/go-auth/internal/model"
	modelUserRepo "github.com/UraharaKiska/go-auth/internal/repository/user/model"
	modelAccessibleRoleRepo "github.com/UraharaKiska/go-auth/internal/repository/accessibleRole/model"
)

type UserRepository interface {
	Create(ctx context.Context, info *modelUserRepo.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *modelUserRepo.UpdateUserInfo, id int64) error
	Delete(ctx context.Context, id int64) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type AccessibleRepository interface {
	GetEndpointRole(ctx context.Context, endpoint string) (*modelAccessibleRoleRepo.EndpointRole, error)
}

