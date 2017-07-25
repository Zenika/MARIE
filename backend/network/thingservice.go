package network

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/record"
	"github.com/Zenika/MARIE/backend/thing"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
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
		mqttConn.AddSubscription(g.Name)
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
