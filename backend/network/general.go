package network

import (
	"log"

	"github.com/Zenika/MARIE/backend/thing"
)

// Protocol represents a thing with a certain protocol
type Protocol interface {
	Type() string
	Do(mac string, name string, params map[string]interface{})
	Get(id string, name string, macaddress string)
}

var protocols map[string]Protocol

// Init network
func Init() {
	protocols = make(map[string]Protocol)
}

// AddProtocol to map
func AddProtocol(p Protocol) {
	protocols[p.Type()] = p
}

// Do something on all things that match action and room
func Do(thingType string, action string, params map[string]interface{}, location string) (int, error) {
	things, err := thing.ReadActionName(action)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, t := range things {
		if t.Type == thingType {
			if location == "" || t.Location == location {
				sum = sum + 1
				protocols[t.Protocol].Do(t.MacAddress, action, params)
			}
		}
	}
	return sum, nil
}

// Get some value
func Get(id string, name string, location string) (int, error) {
	things, err := thing.ReadGetterName(name)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	sum := 0
	for _, t := range things {
		if location == "" || location == t.Location {
			sum = sum + 1
			protocols[t.Protocol].Get(id, name, t.MacAddress)
		}
	}
	return sum, nil
}
