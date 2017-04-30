package main

import (
	"html/template"
	"net/http"
)

var pageTemplate = template.Must(template.ParseFiles("static/index.html"))

func init() {
	http.HandleFunc("/", pageHandler)
}

func pageHandler(res http.ResponseWriter, _ *http.Request) {
	pageTemplate.Execute(res, nil)
}
