package config

import (
	"github.com/joho/godotenv"
	"time"
)

type HTTPServerConfig interface {
	Address() string
	GetTimeout() time.Duration
	GetIdleTimeout() time.Duration
}

type PGConfig interface {
	DSN() string
}

func LoadConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
