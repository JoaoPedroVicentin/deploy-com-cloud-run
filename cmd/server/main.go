package main

import (
	"net/http"

	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/infra/webservers/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Get("/{cep}", handler.GetTemperatureOfLocation)
	})

	http.ListenAndServe(":8080", r)
}
