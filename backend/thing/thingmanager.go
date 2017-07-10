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

// Read all things in database
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
