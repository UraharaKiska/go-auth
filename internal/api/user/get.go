package user

import (
	"context"

	"github.com/UraharaKiska/go-auth/internal/converter"
	// "github.com/UraharaKiska/go-auth/internal/logger"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
	// "go.uber.org/zap"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	// logger.Info("Getting user...", zap.Int64("id", req.GetId()))
	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.GetResponse{
		Info: converter.ToUserSecureFromService(userObj),
	}, nil

}
