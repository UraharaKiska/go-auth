package auth

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-auth/internal/converter"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("API - GET")
	id, err := i.authService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{
		Id: id,
	}, nil

}
