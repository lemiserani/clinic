package main

import (
	"log"
	"os"

	mongo "gopkg.in/mgo.v2"
)

//GetSession creates a new mongo session and panics if connection error occurs
func GetSession() *mongo.Session {
	mongodb := os.Getenv("MONGODB_URI")
	s, _ := mongo.Dial(mongodb)

	if mongodb == "" {
		log.Fatal("$MONGODB_URI must be set")
	}
	// Deliver session
	return s
}
