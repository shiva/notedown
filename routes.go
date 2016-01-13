package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route {"Index", "GET", "/", Index},
    Route {"Add", "GET", "/notes/add", Add},
    Route {"Remove", "GET", "/notes/remove/{noteId}", Remove},
    Route {"List", "GET", "/notes/list", List},
}
