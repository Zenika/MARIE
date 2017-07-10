package thing

import (
	"log"

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
