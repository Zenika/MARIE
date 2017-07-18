package network

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/mqtt"
	"github.com/Zenika/MARIE/backend/record"
	"github.com/Zenika/MARIE/backend/thing"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

// Post handle the post request
func Post(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var t thing.Thing
	for {
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	t.ID = bson.NewObjectId()
	thing.Create(t)

	for _, g := range t.Getters {
		mqtt.AddSubscription(g.Name)
	}
}

// GetAll things and send it
func GetAll(w http.ResponseWriter, r *http.Request) {
	things := thing.ReadAll()

	res, err := json.Marshal(things)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}

// Remove thing from the database
func Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	record.DeleteThingID(bson.ObjectIdHex(vars["id"]))

	thing.Delete(bson.ObjectIdHex(vars["id"]))
}
