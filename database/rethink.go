package database

import (
	"log"
	r "gopkg.in/gorethink/gorethink.v4"
)

var masterSession *r.Session

var connectOpts = r.ConnectOpts{
	Address: "localhost:28015",
	Database: "test",
	MaxOpen: 40,
}

func init() {
	session, err := r.Connect(connectOpts)
	if err != nil {
		log.Fatalln("[FATAL] ", err.Error())
	}
	masterSession = session
}

//GetSession - Clones the master session to derive an additional session
func GetSession() *r.Session {
	return masterSession
}