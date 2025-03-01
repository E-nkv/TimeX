package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
)

type app struct {
	Service Service
}

func NewApp(db *pgx.Conn) *app {
	return &app{
		Service: Service{
			Repo: Repo{DB: db},
		},
	}
}

func (app app) SetupRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello world"
		writeResp(w, http.StatusOK, msg, "data")
	})
	r.Route("/categories", func(r chi.Router) {
		r.Post("/", app.HandleCreateCategory)
		r.Get("/", app.HandleGetCategories)
	})

	r.Route("/sessions", func(r chi.Router) {
		r.Get("/history", app.HandleGetHistory)

		r.Post("/", app.HandleCreateSession)
		r.Get("/{id}", app.HandleGetSession)
		r.Delete("/{id}", app.HandleDeleteSession)

	})

	return r
}
