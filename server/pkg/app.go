package app

//create a api with go and chi server

import (
	// "fmt"
	// "log"
	"net/http"
	// "os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const (
	// Port is the port that the server will listen on
	Port = "8080"

	// Host is the host that the server will listen on
	Host = "localhost"

	// Name is the name of the API
	Name = "Golang API Rest"
)

// App is the application
type App struct {
	Router *chi.Mux
	httpClient *http.Client
	// dependencies *dependencies
}

func (a *App) setupInfraestructure() {
	a.httpClient = &http.Client{}
}

func (a *App) setupServer() {
	a.Router = chi.NewRouter()
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.URLFormat)
	a.Router.Use(render.SetContentType(render.ContentTypeJSON))
}



