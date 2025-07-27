package auth

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-auth/internal/converter"
	"github.com/UraharaKiska/go-auth/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.UserUpdateInfo, id int64) error {
	log.Printf("SERVICE - UPDATE")
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.userRepository.Update(ctx, converter.ToUserUpdateInfoFromService(info), id)
		if errTx != nil {
			return errTx
		}
		return nil
		})
	if err != nil {
		return err
	}
	return nil
}
