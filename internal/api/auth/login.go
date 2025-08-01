package auth

import (
	"context"
	"log"

	// "github.com/UraharaKiska/go-auth/internal/converter"
	"github.com/UraharaKiska/go-auth/internal/converter"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	log.Printf("API - GET")
	token, err := i.authService.Login(ctx, converter.ToUserInfoFromDescLoginRequest(req))
	if err != nil {
		return nil, err
	}
	return &desc.LoginResponse{
		RefreshToken: token,
	}, nil

}

