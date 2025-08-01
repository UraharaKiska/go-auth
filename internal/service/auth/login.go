package user

import (
	"context"

	// "go/token"

	"github.com/UraharaKiska/go-auth/internal/model"
	// "fmt"
	// "log"
	"github.com/UraharaKiska/go-auth/internal/utils"
	// "github.com/UraharaKiska/go-auth/internal/model"
	// "github.com/UraharaKiska/go-auth/internal/repository/user/converter"
)


func (s *serv) Login(ctx context.Context, user *model.UserInfo) (string, error) {
	// get user
	userOdj, err := s.userRepository.GetByEmail(ctx, user.Email)
	// check password
	if err != nil {
		return "", err
	}
	err = utils.CheckPassword(user.Password, userOdj.Info.Password)
	if err != nil {
		return "", err
	}
	userBaseInfo := model.UserBaseInfo{
		Email: userOdj.Info.Email,
		Role: userOdj.Info.Role,
	}
	token, err := utils.GenerateToken(
		userBaseInfo, 
		s.authConfig.RefreshTokenESecretKey(),
		s.authConfig.RefreshTokenExpiration(),
	)
	if err != nil {
		return "", err
	}
	return token, nil
}