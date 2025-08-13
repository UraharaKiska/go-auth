package user

import (
	"context"
	"fmt"
	"log"
	"github.com/UraharaKiska/go-auth/internal/utils"	
	"github.com/UraharaKiska/go-auth/internal/model"
	"github.com/UraharaKiska/go-auth/internal/repository/user/converter"
)

func (s *serv) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	log.Printf("SERVICE - CREATE")
	if info.Password != info.PasswordConfirm {
		return 0, fmt.Errorf("passwords don't match")
	}
	hashPassword, err := utils.HashPassword(info.Password)
	if err != nil {
		return 0, err
	}
	info.Password = hashPassword
	var id int64
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, converter.ToUserInfoFromService(info))
		if errTx != nil {
			return errTx
		}
		_, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}
		return nil
	})

	if err != nil {
		return 0, err
	}
	return id, nil
}
