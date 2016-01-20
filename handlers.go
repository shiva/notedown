package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request) {
    List(w, r)
}

func Add(w http.ResponseWriter, r *http.Request) {
    var note Note
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &note); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

	u, err := repo.FindUser("shiva")
    if err != nil { panic(err) }

    n := repo.InsertNote(u, note)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(n); err != nil {
        panic(err)
    }
}

func Find(w http.ResponseWriter, r *http.Request) {
}

func Remove(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Remove")
}

func List(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

	u, err := repo.FindUser("shiva")
    if err != nil { panic(err) }
	
    notes, err := repo.ListAllNotes(u)
    if err != nil { panic(err) }

    err = json.NewEncoder(w).Encode(notes)
    if err != nil {
        panic(err)
    }
}
