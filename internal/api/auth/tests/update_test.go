package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/UraharaKiska/go-auth/internal/api/auth"
	"github.com/UraharaKiska/go-auth/internal/model"
	"github.com/UraharaKiska/go-auth/internal/service"
	serviceMocks "github.com/UraharaKiska/go-auth/internal/service/mock"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id               = gofakeit.Int64()
		name             = gofakeit.Animal()
		email            = gofakeit.Animal()

		serviceErr = fmt.Errorf("service error")

		validReq = &desc.UpdateRequest{
			Id: id,
			Info: &desc.UpdateUserInfo{
				Name:  wrapperspb.String(name),
				Email: wrapperspb.String(email),

			},
		}

		// partialUpdateNameOnly = &desc.UpdateRequest{
		// 	Id: id,
		// 	Info: &desc.UpdateUserInfo{
		// 		Name: wrapperspb.String(name),
		// 		// Email omitted (nil)
		// 	},
   		// }	

		userUpdateServiceAllFields = &model.UserUpdateInfo{
			Name: model.OptionString{Value: name, Valid: true},
			Email: model.OptionString{Value: email, Valid: true},
		}

		// partialUpdateEmailOnly = &desc.UpdateRequest{
		// 	Id: id,
		// 	Info: &desc.UpdateUserInfo{
		// 		// Name omitted (nil)
		// 		Email: wrapperspb.String(email),
		// 	},
    	// }

		// emptyUpdate = &desc.UpdateRequest{
		// 	Id: id,
		// 	Info: &desc.UpdateUserInfo{}, // Both fields nil
    	// }


		res = &emptypb.Empty{}
	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case all fields",
			args: args{
				ctx: ctx,
				req: validReq,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, userUpdateServiceAllFields, id).Return(nil)
				return mock
			},
		},
		{
			name: "success error case",
			args: args{
				ctx: ctx,
				req: validReq,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, userUpdateServiceAllFields, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewImplementation(authServiceMock)

			updateResp, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, updateResp)
		})
	}

}
