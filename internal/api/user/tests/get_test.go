package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/UraharaKiska/go-auth/internal/api/auth"
	"github.com/UraharaKiska/go-auth/internal/model"
	"github.com/UraharaKiska/go-auth/internal/service"
	serviceMocks "github.com/UraharaKiska/go-auth/internal/service/mock"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc = minimock.NewController(t)

		id = gofakeit.Int64()
		name             = gofakeit.Animal()
		email            = gofakeit.Animal()
		password         = gofakeit.Animal()
		password_confirm = gofakeit.Animal()
		role             = 1
		created_at = gofakeit.Date()
		updated_at = gofakeit.Date()
		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		res = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name: name,
					Email: email,
					Password: password,
					PasswordConfirm: password_confirm,
					Role: desc.EnumRole(role),
				},
				CreatedAt: timestamppb.New(created_at),
				UpdatedAt: timestamppb.New(updated_at),
			},
		}
		resWithoutUpdate = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name: name,
					Email: email,
					Password: password,
					PasswordConfirm: password_confirm,
					Role: desc.EnumRole(role),
				},
				CreatedAt: timestamppb.New(created_at),
				UpdatedAt: nil,
			},
		}


		userWithoutUpdate = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name: name,
				Email: email,
				Password: password,
				PasswordConfirm: password_confirm,
				Role: desc.EnumRole(role).String(),
		},
			CreatedAt: created_at,
			UpdatedAt: sql.NullTime{time.Time{}, false},
			
		}

		user = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name: name,
				Email: email,
				Password: password,
				PasswordConfirm: password_confirm,
				Role: desc.EnumRole(role).String(),
		},
			CreatedAt: created_at,
			UpdatedAt: sql.NullTime{updated_at, true},
			
		}
	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case with update_at",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(user, nil)
				return mock
			},
		},
		{
			name: "success case without update_at",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: resWithoutUpdate,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(userWithoutUpdate, nil)
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
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
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

			resGet, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resGet)
		})
	}

}
