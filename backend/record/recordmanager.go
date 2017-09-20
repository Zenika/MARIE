package record

import (
	"log"
	"time"

	"github.com/Zenika/MARIE/backend/thing"
	"github.com/Zenika/MARIE/backend/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ReadAll(id bson.ObjectId, getter string) ([]Record, error) {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	r := []Record{}

	err := c.Find(bson.M{"thing_id": id, "name": getter}).All(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// MeanLast of every things for a specific parameter
func MeanLast(name string, l string) (float64, error) {
	things, err := thing.ReadGetterName(name)

	if err != nil {
		return 0, err
	}
	c, s := utils.Database(CollectionName)
	defer s.Close()

	r := Record{}
	sum := 0.0
	n := 0.0

	for _, t := range things {
		err := c.Find(bson.M{"thing_id": t.ID}).Sort("-date").One(&r)

		if err == mgo.ErrNotFound {
			continue
		} else if err != nil {
			log.Println(err)
			continue
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
	return sum / n, nil
}

// Save save a thing record to the database with verification
func (r Record) Save() error {
	c, s := utils.Database(CollectionName)
	defer s.Close()

	_, err := thing.Read(r.ThingID)

	if err != nil {
		return err
	}

	r.Date = time.Now()
	err = c.Insert(&r)

	return err
}

// DeleteThingID delete the records that have a specific thing id
func DeleteThingID(id bson.ObjectId) error {
	c, s := utils.Database(CollectionName)
	defer s.Close()
	log.Println("Delete")
	_, err := c.RemoveAll(bson.M{"thing_id": id})
	return err
}
