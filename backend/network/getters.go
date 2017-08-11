package network

// GetAll get value from all things
func GetAll(id string, getter string) {
	for _, p := range protocols {
		p.GetAll(id, getter)
	}
}

// GetLocation get value from all things in a location
func GetLocation(id string, location string, getter string) {
	for _, p := range protocols {
		p.GetLocation(id, location, getter)
	}
}

// GetMacAddress get value from a specific thing
func GetMacAddress(id string, macaddress string, getter string) {
	for _, p := range protocols {
		p.GetMacAddress(id, macaddress, getter)
	}
}
