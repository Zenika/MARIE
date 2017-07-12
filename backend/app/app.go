package app

import (
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/network"
	"github.com/Zenika/MARIE/backend/thing"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// App represents the MARIE application
type App struct {
	Router http.Handler
}

// Initialize the application
func (a *App) Initialize() {
	a.initializeRoutes()
	network.InitMQTT()
}

// Run the application
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	// Allow every origin
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	// Create router with his routes
	r := mux.NewRouter()

	// Websockets
	r.HandleFunc("/ws", network.Handle)

	// MARIE api
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/things", thing.Post).Methods("POST")
	s.HandleFunc("/things", thing.GetAll).Methods("GET")

	a.Router = c.Handler(r)
}
