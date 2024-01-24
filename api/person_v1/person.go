package person_v1

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const basePath = "/api/v1/persons"

func RegisterRoutes(_ context.Context, router *chi.Mux) {

	router.Get(basePath+"/", func(writer http.ResponseWriter, request *http.Request) {})

	router.Get(basePath+"/{id}", func(writer http.ResponseWriter, request *http.Request) {})

	router.Post(basePath+"/", func(writer http.ResponseWriter, request *http.Request) {})

	router.Patch(basePath+"/:{id}", func(writer http.ResponseWriter, request *http.Request) {})
}
