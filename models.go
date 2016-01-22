package main

import (
    "time"
    "gopkg.in/mgo.v2/bson"

)

type Note struct {
    Id        bson.ObjectId `json:"id" bson:"_id"`
    Text      string        `json:"text" bson:"text"`
    UserId    bson.ObjectId `json:"userid" bson:"userid"`
    CreatedAt time.Time     `json:"created-at" bson:"created-at"`
}

type User struct {
    Id        bson.ObjectId `json:"id" bson:"_id"`
    UserId    string        `json:"userid" bson:"userid"`
    FirstName string        `json:"firstname" bson:"firstname"`
    LastName  string        `json:"lastname" bson:"lastname"`
    Notes     []bson.ObjectId `json:"notes" bson:"notes"`
}
