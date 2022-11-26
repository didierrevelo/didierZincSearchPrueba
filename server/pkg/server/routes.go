package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Routes is the routes of the API
func (a *Server) SetupRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "welcome to didierZincSearchPrueba"})
	})
	a.Router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"status": "ok"})
		render.Status(r, http.StatusOK)
	})

	a.Router.Mount("/indexer", indexerRouter(a))
	a.Router.Mount("/mail", mailRouter(a))	
}

func indexerRouter(a *Server) chi.Router {
	r := chi.NewRouter()
	r.Post("/mails", a.dependencies.indexerHandler.IndexEmails)
	
	return r
}

func mailRouter(a *Server) chi.Router {
	r := chi.NewRouter()
	r.Post("/search", a.dependencies.mailHandler.SearchIntoEmail)

	r.Route("/users", func(r chi.Router){
		r.Get("/", a.dependencies.mailHandler.GetUsers)
		r.Get("/{userID}", a.dependencies.mailHandler.GetEmails)
	})

	return r
}
