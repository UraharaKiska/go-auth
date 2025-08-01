package auth

import (
	"context"

	// "github.com/UraharaKiska/go-auth/internal/converter"


	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"

	
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	
	token, err := i.authService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}
	return &desc.GetRefreshTokenResponse{
		RefreshToken: token,
	}, nil

}