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
	return request(res)
}

func request(res apiaigo.Result) map[string]interface{} {
	// If the user wants to Get some data
	if res.Metadata.IntentName == "Get" {
		id := uuid.NewV4()

		if res.Parameters["location"] == "" {
			network.GetAll(id.String(), res.Parameters["variable-name"])
		} else {
			network.GetLocation(id.String(), res.Parameters["location"], res.Parameters["variable-name"])
		}
		// count, err := network.Get(id.String(), res.Parameters["variable-name"], res.Parameters["location"])
		// if err != nil {
		// 	return map[string]interface{}{"error": err.Error()}
		// }
		return map[string]interface{}{
			"executing": id,
			"message":   res.Fulfillment.Speech,
		}
	}

	// If the user wants to Do something
	if res.Metadata.IntentName == "Do" {
		if res.Parameters["location"] == "" {
			network.DoAll(res.Parameters["thing"], res.Parameters["action"], nil)
		} else {
			network.DoLocation(res.Parameters["thing"], res.Parameters["location"], res.Parameters["action"], nil)
		}
		return map[string]interface{}{
			"doing":   res.Parameters["action"],
			"on":      res.Parameters["thing"],
			"in":      res.Parameters["location"],
			"message": res.Fulfillment.Speech,
		}
	}
	return map[string]interface{}{"message": res.Fulfillment.Speech}
}
