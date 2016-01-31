package main

import (
    "net/http"
    "html/template"
    "path/filepath"
    "os"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    cwd, _ := os.Getwd()
    t, err := template.ParseFiles(filepath.Join( cwd,  "templates/" + tmpl + ".html"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
