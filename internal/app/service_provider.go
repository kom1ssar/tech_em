package app

import (
	"context"
	"github.com/kom1ssar/go_common/pkg/closer"
	"github.com/kom1ssar/go_common/pkg/db"
	"github.com/kom1ssar/go_common/pkg/db/pg"
	"github.com/kom1ssar/tech_em/internal/api"
	"github.com/kom1ssar/tech_em/internal/api/person_v1"
	"github.com/kom1ssar/tech_em/internal/config"
	"github.com/kom1ssar/tech_em/internal/config/env"
	"log"
)

type serviceProvider struct {
	httpServerConfig config.HTTPServerConfig
	pgConfig         config.PGConfig

	dbClient db.Client

	personV1Implementation api.PersonV1Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}

}

func (s *serviceProvider) PgConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("eror pg_config init %+v", err)
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPServerConfig {
	if s.httpServerConfig == nil {
		cfg, err := env.NewHTTPServerConfig()
		if err != nil {
			log.Fatalf("eror http_config init %+v", err)
		}
		s.httpServerConfig = cfg
	}
	return s.httpServerConfig
}

func (s *serviceProvider) DbClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("eror pg_client init %+v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("eror pg_client init ping %+v", err.Error())
		}

		closer.Add(func() error {
			cl.Close()
			return nil
		})
		s.dbClient = cl
	}
	return s.dbClient
}

func (s *serviceProvider) PersonImplementationV1(_ context.Context) api.PersonV1Implementation {
	if s.personV1Implementation == nil {
		s.personV1Implementation = person_v1.NewImplementation()
	}

	return s.personV1Implementation
}
