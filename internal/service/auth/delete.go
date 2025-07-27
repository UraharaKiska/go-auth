package auth

import (
	"context"
	"log"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	log.Printf("SERVICE - DELETE")
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.userRepository.Delete(ctx, id)
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
