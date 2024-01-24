package person_v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/kom1ssar/tech_em/internal/api"
	"net/http"
)

const basePath = "/api/v1/persons"

func RegisterRoutes(router *chi.Mux, impl api.PersonV1Implementation) {

	router.Get(basePath+"/", func(writer http.ResponseWriter, request *http.Request) {})

	router.Get(basePath+"/{id}", func(writer http.ResponseWriter, request *http.Request) {})

	router.Post(basePath+"/", func(writer http.ResponseWriter, request *http.Request) {})

	router.Patch(basePath+"/:{id}", func(writer http.ResponseWriter, request *http.Request) {})
}
