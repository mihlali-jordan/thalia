package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Route("/movies", func(r chi.Router) {
			r.Post("/", app.createMovieHandler)
			r.Get("/{id}", app.getMovieHandler)
		})
	})

	return router
}
