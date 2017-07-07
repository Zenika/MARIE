package apiai

import (
	"fmt"

	"github.com/Zenika/MARIE/backend/config"

	apiaigo "github.com/kamalpy/apiai-go"
)

// Analyze the text request and return the results from apiai
func Analyze(req string) apiaigo.Result {
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
