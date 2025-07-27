package tests

import (
	"context"
	"database/sql"
	"fmt"

	// "fmt"
	"testing"

	"github.com/UraharaKiska/go-auth/internal/client/db"
	// "github.com/UraharaKiska/go-auth/internal/client/db/mock"
	txMock "github.com/UraharaKiska/go-auth/internal/client/db/mock"
	"github.com/UraharaKiska/go-auth/internal/model"
	"github.com/UraharaKiska/go-auth/internal/repository"
	repositoryMocks "github.com/UraharaKiska/go-auth/internal/repository/mock"
	modelRepo "github.com/UraharaKiska/go-auth/internal/repository/user/model"
	"github.com/UraharaKiska/go-auth/internal/service/auth"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.UserUpdateInfo
		id int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id               = gofakeit.Int64()
		name             = gofakeit.Animal()
		email            = gofakeit.Animal()

		// serviceErr = fmt.Errorf("service error")
		transactionErr = fmt.Errorf("transaction error")

		fullReq = &model.UserUpdateInfo{
			Name: model.OptionString{
				Value: name,
				Valid: true,
			},
			Email: model.OptionString{
				Value: email,
				Valid: true,
			},
		}

		fullReqRepa = &modelRepo.UpdateUserInfo{
			Name:  sql.NullString{
				String: name,
				Valid: true,
			},
			Email: sql.NullString{
				String: email,
				Valid: true,
			},
		}

	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		err             error
		userRepositoryMock userRepositoryMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: fullReq,
				id: id,
			},
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, fullReqRepa, id).Return(nil)
				// mock.GetMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
			
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "transaction error case",
			args: args{
				ctx: ctx,
				req: fullReq,
				id: id,
			},
			err:  transactionErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, fullReqRepa, id).Return(transactionErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)
			service := auth.NewService(userRepositoryMock, txManagerMock)
			err := service.Update(tt.args.ctx, tt.args.req, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}

}
