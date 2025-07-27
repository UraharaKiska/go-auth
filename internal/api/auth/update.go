package auth

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-auth/internal/converter"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("API - GET")
	err := i.authService.Update(ctx, converter.ToUserUpdateInfoFromDesc(req.GetInfo()), req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil

}
