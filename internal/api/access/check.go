package access

import (
	"context"

	// "github.com/UraharaKiska/go-auth/internal/converter"
	desc "github.com/UraharaKiska/go-auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}