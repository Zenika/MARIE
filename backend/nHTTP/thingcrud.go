package nHTTP

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/nMQTT"
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
		nMQTT.GetConnection().AddGetSubscription(g.Name)
	}

	res, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
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

	err = t.Update()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Remove thing from the database
func Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, err := thing.Read(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		log.Println(err)
		return
	}
	record.DeleteThingID(bson.ObjectIdHex(vars["id"]))
	t.Delete()
}
