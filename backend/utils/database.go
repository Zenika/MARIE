package utils

import (
	"log"

	"github.com/Zenika/MARIE/backend/config"

	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session

// InitDatabase creates the mongodb session and returns the singleton
func InitDatabase() {
	if dbSession == nil {
		log.Println("Creating database session")
		cfg := config.Load()

		// Connection
		dbS, err := mgo.Dial(cfg.DbURL)
		if err != nil {
			log.Fatal(err)
		}

		// Assignment so that dbSession is used
		dbSession = dbS
		log.Println("Database session created ")
	}
}

// Database returns the database connection
func Database(collectionName string) (*mgo.Collection, *mgo.Session) {
	cfg := config.Load()

	s := dbSession.Copy()

	return s.DB(cfg.DbName).C(collectionName), s
}
