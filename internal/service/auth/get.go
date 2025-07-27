package auth

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	log.Printf("SERVICE - GET")
	var user *model.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		user, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
