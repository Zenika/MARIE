package nHTTP

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Zenika/MARIE/backend/nMQTT"
	"github.com/Zenika/MARIE/backend/nWS"
	"github.com/Zenika/MARIE/backend/network"
	"github.com/Zenika/MARIE/backend/record"
	"github.com/Zenika/MARIE/backend/thing"
	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Post handle the post request
func Post(w http.ResponseWriter, r *http.Request) {
	t, err := parseThing(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	t.ID = bson.NewObjectId()
	thing.Create(t)

	for _, g := range t.Getters {
		nMQTT.GetConnection().AddSubscription(g.Name)
	}
}

// GetThing get a thing
func GetThing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Mongo Id", http.StatusBadRequest)
		return
	}
	t, err := thing.Read(bson.ObjectIdHex(vars["id"]))

	if err == mgo.ErrNotFound {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

// GetAll things and send it
func GetAll(w http.ResponseWriter, r *http.Request) {
	things, err := thing.ReadAll()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(things)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

// Update a thing
func Update(w http.ResponseWriter, r *http.Request) {
	t, err := parseThing(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = thing.Update(t)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Remove thing from the database
func Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	record.DeleteThingID(bson.ObjectIdHex(vars["id"]))

	thing.Delete(bson.ObjectIdHex(vars["id"]))
}

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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	err = thing.AddAction(t)
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

	err = thing.AddGetter(t)
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

// Do something on a precise thing
func Do(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var dr thing.DoRequest
	for {
		if err := dec.Decode(&dr); err == io.EOF {
			break
		}
	}
	network.DoUnique(dr.Protocol, dr.MacAddress, dr.Name, nil)
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

	res, err = json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	// Broadcast the thing creation
	nWS.Broadcast(res)
}
