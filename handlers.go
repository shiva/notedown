package main

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome!")
    notes := Notes {
       Note{Name: "Write presentation", CreatedAt: time.Now()},
       Note{Name: "Host meetup", CreatedAt: time.Now()},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(notes); err != nil {
        panic(err)
    }
}

func Add(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Add")
}

func Remove(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Remove")
}

func List(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "List of notes ")
}
