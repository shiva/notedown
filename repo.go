package main

import (
    "fmt"
    "time"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type MongoProperties struct {
    Uri      string
    User     string
    Password string
}

type Repo struct {
    Session *mgo.Session
    Mongo   MongoProperties
}


func (r Repo) InitDBSession() *mgo.Session {
    mongoUrl := fmt.Sprintf("mongodb://%s:%s@%s/notedown",
                            r.Mongo.User,
                            r.Mongo.Password,
                            r.Mongo.Uri)

    if len(opts.Verbose) >= 1 {
        fmt.Printf("Calling Init DB session. Mongo URL: %s\n", mongoUrl)
    }

    s, err := mgo.Dial(mongoUrl)
    if err != nil {
        fmt.Printf("Cannot connect to Mongo URL: %s\n", mongoUrl)
        panic(err)
    }

    return s
}

func (r *Repo) Init(mongo MongoProperties) {
    if len(opts.Verbose) >= 1 {
        fmt.Printf("Mongo URL: %s\n", mongo)
    }

    r.Mongo = mongo
}

func (r *Repo) GetSession() *mgo.Session {
    if r.Session == nil {
        r.Session = r.InitDBSession()
    }

    return r.Session
}

func (r *Repo) FindNote(id bson.ObjectId) (Note, error) {

    n := Note{}
    err := r.GetSession().DB("notedown").C("notes").FindId(id).One(&n)
    // return empty Note if not found
    return n, err
}

func (r *Repo) InsertNote(t Note) Note {
    t.Id = bson.NewObjectId()
    if t.CreatedAt.IsZero() {
        t.CreatedAt = time.Now()
    }

    r.GetSession().DB("notedown").C("notes").Insert(t)

    return t
}

func (r *Repo) DeleteNote (id bson.ObjectId) error {
    return r.GetSession().DB("notedown").C("notes").RemoveId(id)
}

func (r *Repo) ListAllNotes() ([]Note, error) {
    var notes []Note

    err := r.GetSession().DB("notedown").C("notes").Find(nil).All(&notes)
    return notes, err
}
