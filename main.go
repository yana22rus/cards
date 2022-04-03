package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	http.HandleFunc("/", create_card)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8000", nil)

}

type Card struct {
	Title       string
	Description string
	Class       []string
	Type        string
	Rare        string
	Mana        int
	Attack      int
	Life        int
	Keyword     string
}

type card struct {
	Class1 []string
}

func create_card(w http.ResponseWriter, r *http.Request) {

	aaa := card{[]string{"s", "a", "w"}}

	xp := []card{aaa}

	tpl.ExecuteTemplate(w, "index.gohtml", xp)

}
