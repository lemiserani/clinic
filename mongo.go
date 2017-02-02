package main

import mongo "gopkg.in/mgo.v2"

//GetSession creates a new mongo session and panics if connection error occurs
func GetSession() *mongo.Session {
	// Connect to our local mongo
	s, err := mongo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
