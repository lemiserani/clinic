package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Closure is objeto about cust
type Closure struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Clinic  string        `json:"clinic" bson:"clinic"`
	Name    string        `json:"name" bson:"name"`
	Value   float32       `json:"value" bson:"value"`
	Kind    string        `json:"kind" bson:"kind"`
	Day     int           `json:"day" bson:"day"`
	Month   int           `json:"month" bson:"month"`
	Year    int           `json:"year" bson:"year"`
	Created time.Time     `json:"created" bson:"created" orm:"auto_now_add;type(datetime)"`
}

//Closures is Array of Closure
type Closures []Closure
