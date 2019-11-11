package models

import (
	"gopkg.in/mgo.v2/bson"
)

//MyPlanet type for usage in database
type MyPlanet struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"name" json:"name"`
	Climate string        `bson:"climate" json:"climate"`
	Terrain string        `bson:"terrain" json:"terrain"`
	Movies  []string      `bson:"movies" json:"movies"`
}
