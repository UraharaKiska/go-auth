package access

import (
	"github.com/UraharaKiska/go-auth/internal/config"
	"github.com/UraharaKiska/go-auth/internal/service"
	desc "github.com/UraharaKiska/go-auth/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
	authConfig config.AUTHConfig
}

func NewImplementation(accessService service.AccessService, authConfig config.AUTHConfig) *Implementation {
	return &Implementation{
		accessService: accessService,
		authConfig: authConfig,
	}
}
