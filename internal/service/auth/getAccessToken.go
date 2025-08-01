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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *serv) GetAccessToken(ctx context.Context, token string) (string, error) {
	claims, err := utils.VerifyToken(token, s.authConfig.RefreshTokenESecretKey())
		if err != nil {
			return "", status.Errorf(codes.Aborted, "invalid refresh token")
		}

	userObj, err := s.userRepository.GetByEmail(ctx, claims.Email)
	if err != nil {
		return "", err
	}
	token, err = utils.GenerateToken(
			model.UserBaseInfo{
				Email: userObj.Info.Email,
				Role: userObj.Info.Role,
			}, 
			s.authConfig.AccessTokenESecretKey(),
			s.authConfig.AccessTokenExpiration(),
		)
	if err != nil {
		return "", err
	}

	return token, nil
}