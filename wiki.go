package main

import (
	"net/http"
	"html/template"
)

type person struct {
	Ident int16
	First string
	Last string
	Commented string
}

var tpl * template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":5200", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		Ident: 1,
		First: "George",
		Last: "Okello",
		Commented: "This is my first post on the Golang Page",

	}

	p2 := person{
		Ident: 2,
		First: "Arthur",
		Last: "Okello",
		Commented: "This is my first post on the Golang Page",

	}

	xp := []person{p1, p2}
	tpl.ExecuteTemplate(w, "index.gohtml", xp)
}