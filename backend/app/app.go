package app

import (
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/network"
	"github.com/Zenika/MARIE/backend/utils"
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
	utils.InitDatabase()

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
	network.StartHub()
	r.HandleFunc("/ws", network.Handle)

	// MARIE api
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/things", network.Post).Methods("POST")
	s.HandleFunc("/things", network.GetAll).Methods("GET")
	s.HandleFunc("/things/{id}", network.GetThing).Methods("GET")
	s.HandleFunc("/things", network.Update).Methods("PUT")
	s.HandleFunc("/things/{id}", network.Remove).Methods("DELETE")
	s.HandleFunc("/things/register", network.Register).Methods("POST")

	a.Router = c.Handler(r)
	log.Println("HTTP and WS servers started")
}
