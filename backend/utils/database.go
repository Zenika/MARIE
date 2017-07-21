package utils

import (
	"log"

	"github.com/Zenika/MARIE/backend/config"

	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session

// GetSession creates the mongodb session and returns the singleton
func GetSession() *mgo.Session {
	if dbSession == nil {
		cfg := config.Load()

		// Connection
		dbS, err := mgo.Dial(cfg.DbURL)
		if err != nil {
			log.Fatal(err)
		}

		// Assignment so that dbSession is used
		dbSession = dbS
		log.Println("MongoDB Session created ")
	}
	return dbSession.Copy()
}
