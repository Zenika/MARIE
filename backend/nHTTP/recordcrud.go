package nHTTP

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/record"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// GetAllRecords get all records of a thing with a getter name
func GetAllRecords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Mongo Id", http.StatusBadRequest)
		return
	}

	records, err := record.ReadAll(bson.ObjectIdHex(vars["id"]), vars["getter"])
	if err != nil {
		log.Println(err)
		return
	}

	res, err := json.Marshal(records)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
