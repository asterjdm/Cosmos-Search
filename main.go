package main

import (
    _"fmt"
    _"html"
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

    log.Fatal(http.ListenAndServe(":5000", nil))

}

func home(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/home/index.html")
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