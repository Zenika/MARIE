package nHTTP

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Zenika/MARIE/backend/nWS"
	"github.com/Zenika/MARIE/backend/network"
	"github.com/Zenika/MARIE/backend/thing"
	uuid "github.com/satori/go.uuid"

	mgo "gopkg.in/mgo.v2"
)

// Register a thing in the database with its MAC address
func Register(w http.ResponseWriter, r *http.Request) {
	t, err := parseThing(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if t.Protocol == "" {
		t.Protocol = "HTTP"
	}
	t.IPAddress = getIPFromRemoteAddr(r.RemoteAddr)
	t, err = thing.Register(t)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	launchBroadcast("register", t)
	w.Write([]byte("OK"))
}

// AddAction to a registered thing
func AddAction(w http.ResponseWriter, r *http.Request) {
	t, err := parseThing(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	old, err := thing.ReadMacAddress(t.MacAddress)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = old.AddAction(t.Actions)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Println("Not Found")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	launchBroadcast("actions", t)
}

// AddGetter to a registered thing
func AddGetter(w http.ResponseWriter, r *http.Request) {
	t, err := parseThing(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	old, err := thing.ReadMacAddress(t.MacAddress)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = old.AddGetter(t.Getters)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Println("Not found")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	launchBroadcast("getters", t)
}

// Get a value from a thing
func Get(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var gr thing.GetRequest
	for {
		if err := dec.Decode(&gr); err == io.EOF {
			break
		}
	}
	id := uuid.NewV4().String()
	network.GetMacAddress(id, gr.MacAddress, gr.Name)
	w.Write([]byte(id))
}

// Do something on a precise thing
func Do(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var dr thing.DoRequest
	for {
		if err := dec.Decode(&dr); err == io.EOF {
			break
		}
	}
	id := uuid.NewV4().String()
	network.DoMacAddress(id, dr.MacAddress, dr.Name, dr.Parameters)
	w.Write([]byte(id))
}

func parseThing(r io.ReadCloser) (thing.Thing, error) {
	dec := json.NewDecoder(r)
	var t thing.Thing
	for {
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			return t, err
		}
	}
	return t, nil
}

func getIPFromRemoteAddr(ra string) string {
	ip := strings.Split(ra, ":")
	return ip[0]
}

func launchBroadcast(topic string, t thing.Thing) {
	// Transform to JSON
	res, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	var msg map[string]interface{}
	json.Unmarshal(res, &msg)
	msg["topic"] = topic

	// Broadcast the thing creation
	nWS.BroadcastJSON(msg)
}
