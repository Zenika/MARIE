package thing

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/config"

	"github.com/Zenika/MARIE/backend/utils"

	"strings"

	"time"

	"gopkg.in/mgo.v2"
)

// Create a new thing and add it to the database
func Create(t Thing) {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(ThingCollectionName)
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

	c := s.DB(cfg.DbName).C(ThingCollectionName)
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

	c := s.DB(cfg.DbName).C(ThingCollectionName)

	res := Thing{}
	err := c.FindId(id).One(&res)
	if err == mgo.ErrNotFound {
		errMess := []string{"Thing not found", id.String()}
		return res, errors.New(strings.Join(errMess, " "))
	} else if err != nil {
		return res, err
	}

	return res, nil
}

// MeanLastRecord of every things for a specific parameter
func MeanLastRecord(name string) float64 {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	// Select all things with this parameter
	cThing := s.DB(cfg.DbName).C(ThingCollectionName)
	things := []Thing{}

	err := cThing.Pipe([]bson.M{{"$match": bson.M{"getters.name": name}}}).All(&things)
	if err != nil {
		log.Fatal(err)
	}

	cRecord := s.DB(cfg.DbName).C(RecordCollectionName)
	r := Record{}
	sum := 0.0
	n := 0.0
	for _, t := range things {
		err = cRecord.Find(bson.M{"thing_id": t.ID}).Sort("-date").One(&r)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + r.Value.(float64)
		n++
	}
	return sum / n
}

// SaveRecord save a thing record to the database with verification
func SaveRecord(r Record) error {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	cThing := s.DB(cfg.DbName).C(ThingCollectionName)

	res := Thing{}
	err := cThing.FindId(r.ThingID).One(&res)
	if err == mgo.ErrNotFound {
		errMess := []string{"Thing not found", r.ThingID.String()}
		return errors.New(strings.Join(errMess, " "))
	} else if err != nil {
		return err
	}

	cRecord := s.DB(cfg.DbName).C(RecordCollectionName)
	r.Date = time.Now()
	err = cRecord.Insert(&r)

	return err
}

func verifyAttributes(r Record, t Thing) bool {
	correct := true

	return correct
}
