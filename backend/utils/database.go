package utils

import (
	"log"

	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session

// GetSession creates the mongodb session and returns the singleton
func GetSession() *mgo.Session {
	if dbSession == nil {
		// Connection
		dbS, err := mgo.Dial("mongodb://localhost")
		if err != nil {
			log.Fatal(err)
		}

		// Assignment so that dbSession is used
		dbSession = dbS
		log.Println("MongoDB Session created ")
	}
	return dbSession.Copy()
}
