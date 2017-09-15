package app

import (
	"log"
	"net/http"
	"time"

	"github.com/Zenika/MARIE/backend/nHTTP"
	"github.com/Zenika/MARIE/backend/nMQTT"
	"github.com/Zenika/MARIE/backend/nWS"
	"github.com/Zenika/MARIE/backend/network"
	"github.com/Zenika/MARIE/backend/thing"
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

	nMQTT.InitMQTT()
	network.Init()
	network.AddProtocol(nMQTT.GetConnection())
	go checkHeartBeat()
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
	nWS.StartHub()
	r.HandleFunc("/ws", nWS.Handle)

	// MARIE api
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/things", nHTTP.Post).Methods("POST")
	s.HandleFunc("/things", nHTTP.GetAll).Methods("GET")
	s.HandleFunc("/things/{id}", nHTTP.GetThing).Methods("GET")
	s.HandleFunc("/things", nHTTP.Update).Methods("PUT")
	s.HandleFunc("/things/{id}", nHTTP.Remove).Methods("DELETE")
	s.HandleFunc("/things/register", nHTTP.Register).Methods("POST")
	s.HandleFunc("/things/actions", nHTTP.AddAction).Methods("POST")
	s.HandleFunc("/things/getters", nHTTP.AddGetter).Methods("POST")
	s.HandleFunc("/things/do", nHTTP.Do).Methods("POST")
	s.HandleFunc("/things/get", nHTTP.Get).Methods("POST")

	a.Router = c.Handler(r)
	log.Println("HTTP and WS servers started")
}

func checkHeartBeat() {
	for {
		things, err := thing.ReadAll()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range things {
			if t.IsOnline() {
				if t.State == true {
					message := make(map[string]interface{})
					message["topic"] = "state-off"
					message["macaddress"] = t.MacAddress
					nWS.BroadcastJSON(message)
					t.SetState(false)
				}
			}
		}
		time.Sleep(15000 * time.Millisecond)
	}
}
