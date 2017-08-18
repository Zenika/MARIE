package network

// DoAll do something on all things
func DoAll(id string, thingType string, action string, params map[string]interface{}) {
	for _, p := range protocols {
		p.DoAll(id, thingType, action, params)
	}
}

// DoLocation do something on a specific location
func DoLocation(id string, thingType string, location string, action string, params map[string]interface{}) {
	for _, p := range protocols {
		p.DoLocation(id, thingType, location, action, params)
	}
}

// DoMacAddress do something with a mac address
func DoMacAddress(id string, macaddress string, action string, params map[string]interface{}) {
	for _, p := range protocols {
		p.DoMacAddress(id, macaddress, action, params)
	}
}
