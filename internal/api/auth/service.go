package auth

import (
	"github.com/UraharaKiska/go-auth/internal/config"
	"github.com/UraharaKiska/go-auth/internal/service"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
	authConfig config.AUTHConfig
}

func NewImplementation(authService service.AuthService, authConfig config.AUTHConfig) *Implementation {
	return &Implementation{
		authService: authService,
		authConfig: authConfig,
	}
}
