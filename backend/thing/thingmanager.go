package thing

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/utils"
)

// Create a new thing and add it to the database
func Create(t Thing) error {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	return c.Insert(t)
}

// Register a new Thing in the base with mac address
func Register(t Thing) (Thing, error) {
	_, err := ReadMacAddress(t.MacAddress)
	t.ID = bson.NewObjectId()

	if err == mgo.ErrNotFound {
		return t, Create(t)
	}

	return t, errors.New("Thing already registered")
}

// ReadAll things in database
func ReadAll() ([]Thing, error) {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	var things []Thing

	err := c.Find(bson.M{}).All(&things)

	return things, err
}

// Read a thing in the database with its id
func Read(id bson.ObjectId) (Thing, error) {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	res := Thing{}
	err := c.FindId(id).One(&res)

	return res, err
}

// ReadGetterName return things that have a getter with the given name
func ReadGetterName(name string) ([]Thing, error) {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	// Select all things with this parameter
	things := []Thing{}

	err := c.Pipe([]bson.M{{"$match": bson.M{"getters.name": name}}}).All(&things)
	return things, err
}

// ReadMacAddress return thing with mac address
func ReadMacAddress(mac string) (Thing, error) {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	var t Thing

	err := c.Find(bson.M{"macaddress": mac}).One(&t)
	return t, err
}

// ReadActionName return things that have an action with the given name
func ReadActionName(name string) ([]Thing, error) {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	things := []Thing{}

	err := c.Pipe([]bson.M{{"$match": bson.M{"actions.name": name}}}).All(&things)
	return things, err
}

// Update a thing in database
func (t Thing) Update() error {
	c, s := utils.Database(CollectionName)
	defer s.Close()
	return c.Update(bson.M{"_id": t.ID}, bson.M{"getters": t.Getters,
		"actions":       t.Actions,
		"location":      t.Location,
		"protocol":      t.Protocol,
		"name":          t.Name,
		"type":          t.Type,
		"macaddress":    t.MacAddress,
		"ipaddress":     t.IPAddress,
		"state":         t.State,
		"lastheartbeat": t.LastHeartBeat})
}

// Delete the thing from the database
func (t Thing) Delete() error {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	return c.RemoveId(t.ID)
}

// AddAction to an existing thing
func (t Thing) AddAction(actions []Action) error {
	t.Actions = actions
	return t.Update()
}

// AddGetter to an existing thing
func (t Thing) AddGetter(getters []Getter) error {
	t.Getters = getters
	return t.Update()
}

// SetState of a thing
func (t Thing) SetState(state bool) error {
	t.State = state
	return t.Update()
}

// UpdateHeartBeat to current time and update state to true
func (t Thing) UpdateHeartBeat() error {
	t.LastHeartBeat = time.Now()
	return t.Update()
}

// IsOnline checks the state to see if thing is online
func (t Thing) IsOnline() bool {
	return time.Since(t.LastHeartBeat).Seconds() > 15
}
