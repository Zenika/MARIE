package network

// DoAll do something on all things
func DoAll(thingType string, action string, params map[string]interface{}) {
	for _, p := range protocols {
		p.DoAll(thingType, action, params)
	}
}

// DoLocation do something on a specific location
func DoLocation(thingType string, location string, action string, params map[string]interface{}) {
	for _, p := range protocols {
		p.DoLocation(thingType, location, action, params)
	}
}

// DoMacAddress do something with a mac address
func DoMacAddress(macaddress string, action string, params map[string]interface{}) {
	for _, p := range protocols {
		p.DoMacAddress(macaddress, action, params)
	}
}
