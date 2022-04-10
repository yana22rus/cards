package main

import (
	"fmt"
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

type card struct {
	Class_card  []string
	Type_card   []string
	Rare_card   []string
	Attack_card []int
	Mana_card   []int
	Life_card   []int
}

func create_card(w http.ResponseWriter, r *http.Request) {

	mana := []int{}

	for i := 0; i <= 20; i++ {

		mana = append(mana, i)

	}

	attack := []int{}

	for i := 0; i <= 20; i++ {

		attack = append(attack, i)

	}

	life := []int{}

	for i := 1; i <= 20; i++ {

		life = append(life, i)

	}

	class := card{[]string{"Strength", "Intelligence", "Willpower", "Agility", "Endurance", "Netural"}, []string{"Creature", "Action", "Item", "Support"}, []string{"Legendary", "Epic", "Rare", "Common"}, attack, mana, life}

	xp := []card{class}

	tpl.ExecuteTemplate(w, "index.gohtml", xp)

	if r.Method != "POST" {

		input := r.FormValue("title_card")
		fmt.Println(input)

	}

}
