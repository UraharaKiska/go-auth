package user

import (
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/go-auth/internal/repository"
	"github.com/UraharaKiska/go-auth/internal/service"
	"github.com/UraharaKiska/go-auth/internal/config"
)

type serv struct {
	// access_blacklist
	// refresh blacklist
	txManager      db.TxManager
	authConfig     config.AUTHConfig
	accessibleRoleRepo repository.AccessibleRepository
}

func NewService(
	authConfig config.AUTHConfig,
	txManager db.TxManager,
	accessibleRoleRepo repository.AccessibleRepository,
) service.AccessService {
	return &serv{
		authConfig: authConfig,
		txManager:      txManager,
		accessibleRoleRepo: accessibleRoleRepo,
	}
}