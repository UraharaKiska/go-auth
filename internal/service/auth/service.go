package user

import (
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/go-auth/internal/repository"
	"github.com/UraharaKiska/go-auth/internal/service"
	"github.com/UraharaKiska/go-auth/internal/config"
)

type serv struct {
	userRepository repository.UserRepository
	// access_blacklist
	// refresh blacklist
	txManager      db.TxManager
	authConfig     config.AUTHConfig
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
	authConfig config.AUTHConfig,
) service.AuthService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
		authConfig: authConfig,
	}
}