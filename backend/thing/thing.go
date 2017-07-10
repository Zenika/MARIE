package thing

const (
	// CollectionName represents the collection name in the mongo db
	CollectionName = "things"
)

// Thing represents a connected object
type Thing struct {
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Protocol string   `json:"protocol"`
	Actions  []Action `json:"actions"`
	Getters  []Getter `json:"getters"`
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
