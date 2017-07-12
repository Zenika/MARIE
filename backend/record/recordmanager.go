package record

import (
	"log"
	"time"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/thing"
	"github.com/Zenika/MARIE/backend/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MeanLast of every things for a specific parameter
func MeanLast(name string, l string) float64 {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	things := thing.ReadGetterName(name)

	c := s.DB(cfg.DbName).C(CollectionName)
	r := Record{}
	sum := 0.0
	n := 0.0

	for _, t := range things {
		err := c.Find(bson.M{"thing_id": t.ID}).Sort("-date").One(&r)

		if err == mgo.ErrNotFound {
			continue
		} else if err != nil {
			log.Fatal(err)
		}
		if l == "" {
			sum = sum + r.Value.(float64)
			n++
		} else {
			if t.Location == l {
				sum = sum + r.Value.(float64)
				n++
			}
		}
	}
	return sum / n
}

// Save save a thing record to the database with verification
func Save(r Record) error {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	_, err := thing.Read(r.ThingID)

	if err != nil {
		return err
	}

	c := s.DB(cfg.DbName).C(CollectionName)
	r.Date = time.Now()
	err = c.Insert(&r)

	return err
}
