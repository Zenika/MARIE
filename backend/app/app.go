package app

import (
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/mqtt"
	"github.com/Zenika/MARIE/backend/network"
	"github.com/Zenika/MARIE/backend/websocket"
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
	mqtt.Init()
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
	r.HandleFunc("/ws", websocket.Handle)

	// MARIE api
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/things", network.Post).Methods("POST")
	s.HandleFunc("/things", network.GetAll).Methods("GET")
	s.HandleFunc("/things/{id}", network.Remove).Methods("DELETE")

	a.Router = c.Handler(r)
	log.Println("HTTP and WS servers started")
}
