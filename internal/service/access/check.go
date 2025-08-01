package user

import (
	"context"
	"log"
	// "go/token"
	"strings"

	// "github.com/UraharaKiska/go-auth/internal/model"
	// "fmt"
	// "log"
	"github.com/UraharaKiska/go-auth/internal/utils"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	// "github.com/UraharaKiska/go-auth/internal/model"
	// "github.com/UraharaKiska/go-auth/internal/repository/user/converter"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)


func (s *serv) Check(ctx context.Context, endpoint string) (error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md[s.authConfig.AuthHeader()]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], s.authConfig.AuthPrefix()) {
		return errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], s.authConfig.AuthPrefix())
	log.Printf("Token:%v", accessToken)
	claims, err := utils.VerifyToken(accessToken, s.authConfig.AccessTokenESecretKey())
	if err != nil {
		return errors.New("access token is invalid")
	}

	endpointRole, err := s.accessibleRoleRepo.GetEndpointRole(ctx, endpoint)
	if err != nil {
		return err
	}

	if endpointRole == nil {
		return nil
	}

	if endpointRole.Role == claims.Role {
		return nil
	}

	return errors.New("access denied")
}
