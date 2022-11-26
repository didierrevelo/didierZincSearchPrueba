package server

//create a api with go and chi server

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	// "github.com/joho/godotenv"

	zincsearch "github.com/didierrevelo/didierZincSearchPrueba/server/pkg/databaseAdap/zincSearch"
	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/handlers"
	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/service"
)

const (
	// Port is the port that the server will listen on
	Port = "8080"

	// Host is the host that the server will listen on
	Host = "http://localhost:"

	// Name is the name of the API
	Name = "didierZincSearchPrueba"
)

// Server is the Serverlication
type Server struct {
	Router       *chi.Mux
	httpClient   *http.Client
	dependencies *dependencies
}

type dependencies struct {
	mailHandler    *handlers.MailHandler
	indexerHandler *handlers.Indexer
}

func (a *Server) setupInfraestructure() {
	a.httpClient = &http.Client{}
}

func (a *Server) dependenciesSet() {
	zincSearchAdap := zincsearch.NewZincSearchClient(a.httpClient)

	emailService := service.NewEmailService(zincSearchAdap)
	emailHandler := handlers.NewMailHandler(emailService)

	indexerService := service.NewIndexerService(zincSearchAdap)
	indexerHandler := handlers.NewIndexer(indexerService, emailService)

	a.dependencies = &dependencies{
		mailHandler:    emailHandler,
		indexerHandler: indexerHandler,
	}
}

func (a *Server) setupServer() {
	a.Router = chi.NewRouter()
	a.Router.Use(middleware.RequestID)
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.URLFormat)
	a.Router.Use(render.SetContentType(render.ContentTypeJSON))
	
	a.SetupRoutes()
}

// NewServer creates a new Server
func NewServer() *Server {
	server := &Server{}
	server.setupInfraestructure()
	server.dependenciesSet()
	server.setupServer()

	return server
}

// StartServer starts the Serverlication
func (a *Server) StartServer() {
	appName := os.Getenv("SERVER_NAME")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if appName == "" || port == "" {
		port = Port
		appName = Name
		log.Println("port and app name environment variables were not set, default values will be used")
	}
	log.Printf("Server is running on: %s%s\n", host, port)

	if err := http.ListenAndServe(":"+Port, a.Router); err != nil {
		panic(err)
	}
}
