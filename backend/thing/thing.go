package thing

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// ThingCollectionName represents the collection name in the mongo db
	ThingCollectionName = "things"
)

// Thing represents a connected object
type Thing struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	MacAddress string        `json:"macaddress"`
	Type       string        `json:"type"`
	Name       string        `json:"name"`
	Location   string        `json:"location"`
	Protocol   string        `json:"protocol"`
	Actions    []Action      `json:"actions"`
	Getters    []Getter      `json:"getters"`
}

// Action represents what a thing can do
type Action struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}

// Parameter represents what an action needs to be executed
type Parameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Getter represents what information a thing can give
type Getter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
