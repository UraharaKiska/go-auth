package app

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-auth/internal/api/access"
	"github.com/UraharaKiska/go-auth/internal/api/auth"
	"github.com/UraharaKiska/go-auth/internal/api/user"
	"github.com/UraharaKiska/go-auth/internal/config"
	env "github.com/UraharaKiska/go-auth/internal/config/env"
	"github.com/UraharaKiska/go-auth/internal/repository"
	accessibleRoleRepository "github.com/UraharaKiska/go-auth/internal/repository/accessibleRole"
	userRepository "github.com/UraharaKiska/go-auth/internal/repository/user"
	"github.com/UraharaKiska/platform-common/pkg/closer"
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/platform-common/pkg/db/pg"
	"github.com/UraharaKiska/platform-common/pkg/db/transaction"

	"github.com/UraharaKiska/go-auth/internal/service"
	accessService "github.com/UraharaKiska/go-auth/internal/service/access"
	authService "github.com/UraharaKiska/go-auth/internal/service/auth"
	userService "github.com/UraharaKiska/go-auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig
	authConfig config.AUTHConfig
	swaggerConfig config.SWAGGERConfig
	prometheusConfig config.PrometheusConfig


	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository
	accessibleRoleRepository repository.AccessibleRepository

	authService service.AuthService
	userService service.UserService
	accessService service.AccessService

	userImpl *user.Implementation
	authImpl *auth.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.GRPCConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SWAGGERConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		s.swaggerConfig = cfg
	}
	return s.swaggerConfig
}

func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusConfig == nil {
		cfg, err := env.NewPrometheusConfig()
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		s.prometheusConfig = cfg
	}
	return s.prometheusConfig
}

func (s *serviceProvider) AuthConfig() config.AUTHConfig {
	if s.authConfig == nil {
		cfg, err := env.NewAUTHConfig()
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		s.authConfig = cfg
	}
	return s.authConfig
}



func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to init app%v:", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}
	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactorManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) AccessibleRoleRepository(ctx context.Context) repository.AccessibleRepository {
	if s.accessibleRoleRepository == nil {
		s.accessibleRoleRepository = accessibleRoleRepository.NewRepository(s.DBClient(ctx))
	}

	return s.accessibleRoleRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.AuthConfig(),
			s.TxManager(ctx),
			s.AccessibleRoleRepository(ctx),
		)
	}

	return s.accessService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.AuthConfig(),
		)
	}

	return s.authService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx), s.AuthConfig())
	}
	return s.authImpl
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx), s.AuthConfig())
	}
	return s.accessImpl
}

