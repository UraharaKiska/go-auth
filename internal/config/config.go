package config

import (
	"flag"
	"time"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

type GRPCConfig interface {
	Address() string
}

type HTTPConfig interface {
	Address() string
}

type SWAGGERConfig interface {
	Address() string
}

type PrometheusConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type AUTHConfig interface {
	RefreshTokenExpiration() time.Duration 
	AccessTokenExpiration() time.Duration 
	RefreshTokenESecretKey() []byte
	AccessTokenESecretKey() []byte
	AuthHeader() string
	AuthPrefix() string
}


func ParseConfig() string {
	var configPath string
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
	flag.Parse()
	return configPath
}
