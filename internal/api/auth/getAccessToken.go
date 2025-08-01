package auth

import (
	"context"

	// "github.com/UraharaKiska/go-auth/internal/converter"


	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"

	
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	token, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}
	return &desc.GetAccessTokenResponse{
		AccessToken: token,
	}, nil

}