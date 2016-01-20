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
    Route {"Add", "POST", "/notes/add", Add},
    Route {"Remove", "GET", "/notes/remove/{noteId}", Remove},
    Route {"Find", "GET", "/notes/find/{noteId}", Find},
    Route {"List", "GET", "/notes/list", List},
}
