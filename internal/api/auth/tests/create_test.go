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
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id               = gofakeit.Int64()
		name             = gofakeit.Animal()
		email            = gofakeit.Animal()
		password         = gofakeit.Animal()
		password_confirm = gofakeit.Animal()
		role             = 1

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Info: &desc.UserInfo{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password_confirm,
				Role:            desc.EnumRole(role),
			},
		}

		info = &model.UserInfo{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password_confirm,
			Role:            desc.EnumRole(role).String(),
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "success error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
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

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}

}
