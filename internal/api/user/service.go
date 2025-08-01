package user

import (
	"github.com/UraharaKiska/go-auth/internal/service"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
