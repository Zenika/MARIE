package apiai

import (
	"fmt"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/record"

	apiaigo "github.com/kamalpy/apiai-go"
)

// Analyze request and returns JSON
func Analyze(req string) map[string]interface{} {
	res := request(req)
	if res.Metadata.IntentName == "Get" {
		if res.Parameters["room"] != "" {
			return map[string]interface{}{"mean": record.MeanLast(res.Parameters["variable-name"], res.Parameters["room"])}
		}
		return map[string]interface{}{"mean": record.MeanLast(res.Parameters["variable-name"], "")}
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
