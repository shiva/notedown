package main

import (
    "net/http"
)

type RouteType int

const (
    UNSECURED RouteType = 1 + iota
    SECURED
)

var routeTypes = [...]string {
    "unsecured",
    "secured",
}

func (rtype RouteType) String() string {
    return routeTypes[rtype - 1]
}


type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
    rtype       RouteType
}

type Routes []Route

var routes = Routes {
    Route {"Index", "GET", "/", Index, UNSECURED},
    Route {"Add", "POST", "/notes", Add, SECURED},
    Route {"Remove", "DELETE", "/notes/{noteId}", Remove, SECURED},
    Route {"Find", "GET", "/notes/{noteId}", Find, SECURED},
    Route {"List", "GET", "/notes", List, SECURED},
    //Route {"User", "GET", "/user", UserHandler, SECURED},
}
