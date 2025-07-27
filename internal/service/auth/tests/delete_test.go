package tests

import (
	"context"
	"fmt"
	// "fmt"
	"testing"

	"github.com/UraharaKiska/go-auth/internal/client/db"
	// "github.com/UraharaKiska/go-auth/internal/client/db/mock"
	txMock "github.com/UraharaKiska/go-auth/internal/client/db/mock"
	"github.com/UraharaKiska/go-auth/internal/repository"
	repositoryMocks "github.com/UraharaKiska/go-auth/internal/repository/mock"
	"github.com/UraharaKiska/go-auth/internal/service/auth"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id               = gofakeit.Int64()

		// serviceErr = fmt.Errorf("service error")
		transactionErr = fmt.Errorf("transaction error")

	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            int64
		err             error
		userRepositoryMock userRepositoryMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
				req: id,
			},
			want: 0,
			err:  transactionErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(transactionErr)
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

			err := service.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}

}
