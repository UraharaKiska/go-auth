package auth

import (
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/go-auth/internal/repository"
	"github.com/UraharaKiska/go-auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.AuthService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}