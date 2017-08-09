package apiai

import (
	"fmt"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/network"
	uuid "github.com/satori/go.uuid"

	apiaigo "github.com/kamalpy/apiai-go"
)

// Analyze request and return the result
func Analyze(req string) map[string]interface{} {
	cfg := config.Load()

	ai := apiaigo.APIAI{
		AuthToken: cfg.APIAiToken,
		Language:  "fr-FR",
		SessionID: "MARIE",
		Version:   "1",
	}

	resp, err := ai.SendText(req)
	if err != nil {
		fmt.Println(err)
	}
	res := resp.Result
	// If the user wants to Get some data
	if res.Metadata.IntentName == "Get" {
		id := uuid.NewV4()
		count, err := network.Get(id.String(), res.Parameters["variable-name"], res.Parameters["location"])
		if err != nil {
			return map[string]interface{}{"error": err.Error()}
		}
		return map[string]interface{}{
			"executing": id,
			"count":     count,
			"message":   res.Fulfillment.Speech,
		}
	}

	// If the user wants to Do something
	if res.Metadata.IntentName == "Do" {
		count, err := network.Do(res.Parameters["thing"], res.Parameters["action"], nil, res.Parameters["location"])
		if err != nil {
			return map[string]interface{}{"error": err.Error()}
		}
		return map[string]interface{}{
			"doing":   res.Parameters["action"],
			"on":      res.Parameters["thing"],
			"in":      res.Parameters["location"],
			"message": res.Fulfillment.Speech,
			"count":   count,
		}
	}
	return map[string]interface{}{"message": res.Fulfillment.Speech}
}
