package db

import (
	"fmt"

	config "../config"

	"gopkg.in/mgo.v2"
)

//TaskCollection It is used to connect collection
var TaskCollection *mgo.Collection

//Connect Used to establish connection with DB
func Connect() {

	dbSession, err := mgo.Dial(config.URI)
	if err != nil {
		fmt.Println("Unable to connect DB")
	}
	TaskCollection = dbSession.DB("CLI").C("tasks")
}
