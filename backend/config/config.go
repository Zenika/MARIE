package config

import (
	"encoding/json"
	"log"
	"os"
)

var cfg Configuration

// Configuration represent the configuration of the application
type Configuration struct {
	APIAiToken string `json:"apiai-token"`
	DbName     string `json:"database-name"`
	DbURL      string `json:"database-url"`
	MQTTUrl    string `json:"mqtt-url"`
	MQTTId     string `json:"mqtt-id"`
}

// Load the configuration and save it
func Load() Configuration {
	if cfg.APIAiToken != "" {
		return cfg
	}
	conf, err := os.Open("config/config.json")

	if err != nil {
		log.Fatal("config:Load error : ", err)
	}

	decoder := json.NewDecoder(conf)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("config:Load error : ", err)
	}

	return cfg
}
