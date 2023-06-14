package main

import (
	"log"
	"net/http"
	"html/template"
	"example.com/Cosmos-Search/engines"
)

func main() {
	_, err := engines.Search("golang programming language")
	if err != nil {
		log.Fatal(err)
	}
	
	http.HandleFunc("/", home)
	http.HandleFunc("/search/", search)

	// Serve les fichiers statiques depuis le répertoire "static"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/home/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func search(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/search/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Mon titre de page",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}