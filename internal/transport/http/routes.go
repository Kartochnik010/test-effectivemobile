package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		links:
			GET    /people
			DELETE /people/{id}
			PUT    /people/{id}
			POST   /people
		`))
	})
	r.Get("/people", h.SearchPerson)
	r.Get("/people/{id}", h.FindPersonById)
	r.Delete("/people/{id}", h.DeletePersonById)
	r.Put("/people/{id}", h.UpdatePersonById)
	r.Post("/people", h.InsertPerson)

	return r
}
