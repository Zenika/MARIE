package thing

import (
	"encoding/json"
	"log"
	"io"
	"net/http"
)

// Post handle the post request
func Post(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var t Thing
	for {
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	Create(t)
}
