package env

import (
	"errors"
	"github.com/kom1ssar/tech_em/internal/config"
	"os"
)

var _ config.PGConfig = (*pgConfig)(nil)

const (
	pgDSNEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func (p *pgConfig) DSN() string {
	return p.dsn
}

func NewPgConfig() (config.PGConfig, error) {
	dsn := os.Getenv(pgDSNEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg_dsn not found")
	}

	return &pgConfig{dsn: dsn}, nil

}
