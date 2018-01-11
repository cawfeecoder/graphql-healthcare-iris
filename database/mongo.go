package database

import (
	"gopkg.in/mgo.v2"
	"log"
)

var masterSession *mgo.Session

var dialInfo = &mgo.DialInfo{
	Addrs: []string{"localhost:27017"},
	Database: "healthcare",
	Source: "admin",
	Username: "mongoadmin",
	Password: "mongoadmin",
}

func init() {
	session, err := mgo.DialWithInfo(dialInfo); if err != nil {
		log.Fatalln("[FATAL] ", err.Error())
	} else {
		log.Print("[INFO] Successfully connected to database")
	}

	masterSession = session
}

//GetSession - Clones the master session to derive an additional session
func GetSession() *mgo.Session {
	return masterSession.Copy()
}