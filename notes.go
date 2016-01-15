package main

import (
    "time"
    "gopkg.in/mgo.v2/bson"

)

type Note struct {
    Id        bson.ObjectId `json:"id" bson:"_id"`
    Name      string        `json:"name" bson:"name"`
    CreatedAt time.Time     `json:"created-at" bson:"created-at"`
}
