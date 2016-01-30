package main

import (
    "net/http"
    "encoding/base64"
    "os"

    "github.com/gorilla/mux"
    "github.com/auth0/go-jwt-middleware"
    "github.com/codegangsta/negroni"
    "github.com/dgrijalva/jwt-go"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            decoded, err := base64.URLEncoding.DecodeString(os.Getenv("AUTH0_CLIENT_SECRET"))
            if err != nil {
                return nil, err
            }
            return decoded, nil
        },
    })

    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        if route.rtype == SECURED {
            handler = negroni.New(
                negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
                negroni.Wrap(handler),
                )
        }

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}
