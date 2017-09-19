package record

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionName represents the collection name in the mongo db
	CollectionName = "records"
)

// Record represents the record for a thing getter
type Record struct {
	ThingID bson.ObjectId `json:"thing_id" bson:"thing_id"`
	Name    string        `json:"name"`
	Value   interface{}   `json:"value"`
	Date    time.Time     `json:"date"`
}
