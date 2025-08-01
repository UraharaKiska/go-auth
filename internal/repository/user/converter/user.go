package converter

import (
	"github.com/UraharaKiska/go-auth/internal/model"
	modelRepo "github.com/UraharaKiska/go-auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(info modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:            info.Name,
		Email:           info.Email,
		Password:        info.PasswordHash,
		Role:            info.Role,
	}
}

func ToUserInfoFromService(info *model.UserInfo) *modelRepo.UserInfo {
	return &modelRepo.UserInfo{
		Name:            info.Name,
		Email:           info.Email,
		PasswordHash:        info.Password,
		Role:            info.Role,
	} 
}
