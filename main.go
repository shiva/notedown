package main

import (
    "log"
    "net/http"
)

func main() {
    rtr := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", rtr))
}

