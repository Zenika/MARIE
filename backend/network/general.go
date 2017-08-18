package network

// Protocol represents a thing with a certain protocol
type Protocol interface {
	Type() string

	// Actions
	DoAll(id string, thingType string, action string, params map[string]interface{})
	DoLocation(id string, thingType string, location string, action string, params map[string]interface{})
	DoMacAddress(id string, action string, macaddress string, params map[string]interface{})

	GetAll(id string, getter string)
	GetLocation(id string, location string, getter string)
	GetMacAddress(id string, macaddress string, getter string)
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

// Get some value
// func Get(id string, name string, location string) (int, error) {
// 	things, err := thing.ReadGetterName(name)
// 	if err != nil {
// 		log.Println(err)
// 		return 0, err
// 	}
// 	sum := 0
// 	for _, t := range things {
// 		if location == "" || location == t.Location {
// 			sum = sum + 1
// 			protocols[t.Protocol].Get(id, name, t.MacAddress)
// 		}
// 	}
// 	return sum, nil
// }
