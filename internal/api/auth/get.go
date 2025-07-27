package auth

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-auth/internal/converter"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("API - GET")
	userObj, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.GetResponse{
		User: converter.ToUserFromService(userObj),
	}, nil

}
