package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(notes); err != nil {
        panic(err)
    }
}

func Add(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Add")
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

    n := RepoCreateNote(note)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(n); err != nil {
        panic(err)
    }
}

func Remove(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Remove")
}

func List(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "List of notes ")
}
