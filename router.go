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

        handler = Logger(route.HandlerFunc, route.Name)

        if route.rtype == SECURED {
            handler = negroni.New(
                negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
                negroni.Wrap(handler),
                )
        }

        router.Handle(route.Pattern, handler).Name(route.Name).Methods(route.Method)
    }

    router.HandleFunc("/auth", CallbackHandler)
	router.Handle("/user", negroni.New(
		negroni.HandlerFunc(IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(UserHandler)),
	))

    router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

    return router
}
