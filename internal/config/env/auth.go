package config

import (
	"os"
	"strconv"
	"time"

	"github.com/UraharaKiska/go-auth/internal/config"
	"github.com/pkg/errors"
)

const (
	refreshTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION"
	refreshTokenESecretKeyEnvName = "REFRESH_TOKEN_SECRET_KEY"
	accessTokenESecretKeyEnvName = "ACCESS_TOKEN_SECRET_KEY"

)

type authConfig struct {
	refreshTokenExpiration time.Duration
	accessTokenExpiration time.Duration
	refreshTokenESecretKey []byte
	accessTokenESecretKey  []byte
	authHeader string
	authPrefix string
}

func NewAUTHConfig() (config.AUTHConfig, error) {
	refreshTokenExpiration := os.Getenv(refreshTokenExpirationEnvName)
	if len(refreshTokenExpiration) == 0 {
		return nil, errors.New("grpc host not found")
	}
	refreshTokenExpirationInt, err := strconv.Atoi(refreshTokenExpiration)
	if err != nil {
		return nil, errors.New("Cannot convert accessTokenExpiration to int64")
	}

	accessTokenExpiration := os.Getenv(accessTokenExpirationEnvName)
	if len(accessTokenExpiration) == 0 {
		return nil, errors.New("grpc host not found")
	}
	accessTokenExpirationInt, err := strconv.Atoi(accessTokenExpiration)
	if err != nil {
		return nil, errors.New("Cannot convert accessTokenExpiration to int64")
	}

	refreshTokenESecretKey := os.Getenv(refreshTokenESecretKeyEnvName)
	if len(refreshTokenESecretKey) == 0 {
		return nil, errors.New("grpc host not found")
	}

	accessTokenESecretKey := os.Getenv(accessTokenESecretKeyEnvName)
	if len(accessTokenESecretKey) == 0 {
		return nil, errors.New("grpc host not found")
	}

	return &authConfig{
		refreshTokenExpiration: time.Minute * time.Duration(accessTokenExpirationInt),
		accessTokenExpiration: time.Minute * time.Duration(refreshTokenExpirationInt),
		refreshTokenESecretKey:[]byte(refreshTokenESecretKey),
		accessTokenESecretKey: []byte(accessTokenESecretKey),
		authHeader: "authorization",
		authPrefix: "Bearer ",
	}, nil
}

func (a *authConfig) RefreshTokenExpiration() time.Duration {
	return a.refreshTokenExpiration
}

func (a *authConfig) AccessTokenExpiration() time.Duration {
	return a.accessTokenExpiration
}

func (a *authConfig) RefreshTokenESecretKey() []byte {
	return a.refreshTokenESecretKey
}

func (a *authConfig) AccessTokenESecretKey() []byte {
	return a.accessTokenESecretKey
}

func (a *authConfig) AuthHeader() string {
	return a.authHeader
}

func (a *authConfig) AuthPrefix() string{
	return a.authPrefix
}




// func (cfg *httpConfig) Address() string {
// 	return net.JoinHostPort(cfg.host, cfg.port)
// }
