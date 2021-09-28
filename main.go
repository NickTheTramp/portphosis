package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Title"}
	t, _ := template.ParseFiles("templates/index.tmpl.html","templates/base.tmpl.html")
	t.Execute(w, p)
}

func otherHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Other"}
	t, _ := template.ParseFiles("templates/index.tmpl.html","templates/base.tmpl.html")
	t.Execute(w, p)
}


func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/other", otherHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
