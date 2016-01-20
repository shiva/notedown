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
    Info    *mgo.DialInfo
}


func (r Repo) InitDBSession() *mgo.Session {
    s, err := mgo.DialWithInfo(r.Info)
    if err != nil {
        fmt.Printf("Cannot connect to Mongo URL: %s\n", err)
        panic(err)
    }

    return s
}

func (r *Repo) Init(mongo MongoProperties) {
    if len(opts.Verbose) >= 1 {
        fmt.Printf("Mongo URL: %s\n", mongo)
    }

    r.Info = &mgo.DialInfo{
        Addrs: []string{mongo.Uri},
        Timeout:  60 * time.Second,
        Database: "notedown",
        Username: mongo.User,
        Password: mongo.Password,
    }

    r.Session = r.InitDBSession()
}

func (r *Repo) GetSession() *mgo.Session {
    return r.Session
}

func (r *Repo) FindNote(id bson.ObjectId) (Note, error) {

    n := Note{}
    err := r.GetSession().DB("notedown").C("notes").FindId(id).One(&n)
    // return empty Note if not found
    return n, err
}

func (r *Repo) InsertNote(u User, t Note) Note {
    t.Id = bson.NewObjectId()
    if t.CreatedAt.IsZero() {
        t.CreatedAt = time.Now()
    }

    t.UserId = u.Id
    r.GetSession().DB("notedown").C("notes").Insert(t)

    return t
}

func (r *Repo) DeleteNote (id bson.ObjectId) error {
    return r.GetSession().DB("notedown").C("notes").RemoveId(id)
}

func (r *Repo) ListAllNotes(u User) ([]Note, error) {
    var notes []Note

    err := r.GetSession().DB("notedown").C("notes").Find(
        bson.M{"userid": u.Id}).All(&notes)
    return notes, err
}

func (r *Repo) FindUser(name string) (User, error) {
    var user User
    err := r.GetSession().DB("notedown").C("users").Find(bson.M{"userid": name}).One(&user)

    return user, err
}
