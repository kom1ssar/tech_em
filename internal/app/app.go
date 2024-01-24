package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kom1ssar/go_common/pkg/closer"
	"github.com/kom1ssar/tech_em/api/person_v1"
	"github.com/kom1ssar/tech_em/internal/config"
	"log"
	"net/http"
)

type app struct {
	httpServer      *http.Server
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) *app {
	a := &app{}

	err := a.initDeps(ctx)
	if err != nil {
		log.Fatalf("err")
	}

	return a
}

func (a *app) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHttpServer,
	}

	for _, f := range inits {

		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *app) initConfig(_ context.Context) error {
	err := config.LoadConfig("local.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *app) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *app) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runHttpSerer()

}

func (a *app) initHttpRoutesAndMiddleware(ctx context.Context) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	person_v1.RegisterRoutes(router, a.serviceProvider.PersonImplementationV1(ctx))

	return router
}

func (a *app) initHttpServer(ctx context.Context) error {

	handler := a.initHttpRoutesAndMiddleware(ctx)
	a.httpServer = &http.Server{
		Addr:         a.serviceProvider.HTTPConfig().Address(),
		IdleTimeout:  a.serviceProvider.HTTPConfig().GetIdleTimeout(),
		ReadTimeout:  a.serviceProvider.HTTPConfig().GetTimeout(),
		WriteTimeout: a.serviceProvider.HTTPConfig().GetTimeout(),
		Handler:      handler,
	}

	return nil
}

func (a *app) runHttpSerer() error {

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
