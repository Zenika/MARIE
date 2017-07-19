package apiai

import (
	"fmt"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/record"

	"github.com/Zenika/MARIE/backend/network"
	apiaigo "github.com/kamalpy/apiai-go"
)

// Analyze request and returns JSON
func Analyze(req string) map[string]interface{} {
	res := request(req)
	if res.Metadata.IntentName == "Get" {
		if res.Parameters["room"] != "" {
			return map[string]interface{}{"mean": record.MeanLast(res.Parameters["variable-name"], res.Parameters["location"])}
		}
		return map[string]interface{}{"mean": record.MeanLast(res.Parameters["variable-name"], "")}
	}

	if res.Metadata.IntentName == "Do" {
		network.Do(res.Parameters["thing"], res.Parameters["action"], nil, res.Parameters["location"])
		return map[string]interface{}{"doing": res.Parameters["action"], "on": res.Parameters["thing"], "in": res.Parameters["location"]}
	}
	return nil
}

// Request apiai and return the result
func request(req string) apiaigo.Result {
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

	return resp.Result
}
