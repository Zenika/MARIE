package thing

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/config"

	"github.com/Zenika/MARIE/backend/utils"
)

// Create a new thing and add it to the database
func Create(t Thing) {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(CollectionName)
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadAll things in database
func ReadAll() []Thing {
	cfg := config.Load()
	s := utils.GetSession()
	defer s.Close()

	var things []Thing

	c := s.DB(cfg.DbName).C(CollectionName)
	err := c.Find(bson.M{}).All(&things)
	if err != nil {
		log.Fatal(err)
	}

	return things
}

// Read a thing in the database with its id
func Read(id bson.ObjectId) (Thing, error) {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(CollectionName)

	res := Thing{}
	err := c.FindId(id).One(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}

// ReadGetterName return things that have a getter with the given name
func ReadGetterName(name string) []Thing {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	// Select all things with this parameter
	c := s.DB(cfg.DbName).C(CollectionName)
	things := []Thing{}

	err := c.Pipe([]bson.M{{"$match": bson.M{"getters.name": name}}}).All(&things)
	if err != nil {
		log.Fatal(err)
	}
	return things
}
