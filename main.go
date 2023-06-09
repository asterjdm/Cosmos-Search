package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/asterjdm/Cosmos-Search/engines"
)

func main() {

	fmt.Println("Launched on : http://localhost:5000")
	
	http.HandleFunc("/", home)
	http.HandleFunc("/search/", search)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":5000", nil))
	


}

func home(w http.ResponseWriter, r *http.Request) {
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

func search(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	startTime := now.UnixMilli()

	queryParams := r.URL.Query()
	query := queryParams.Get("q")
	pageString := queryParams.Get("page")

	var page int
	var err error
	if pageString != "" {
		page, err = strconv.Atoi(pageString)
		if err != nil {
			page = 0
		}
	}

	var next_page_url string = fmt.Sprintf("/search?q=%s&page=%d", query, page+1)
	results, err := engines.Search(query, &page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/search/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wikiInfo, err := engines.GetWiki(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	FoundInfo, ok := wikiInfo["Found"].(bool)
	if !ok {
		http.Error(w, "Erreur lors de la conversion du champ Found en booléen", http.StatusInternalServerError)
		return
	}

	now = time.Now()
	endTime := now.UnixMilli()
	totalSearchTimeSec := (float64(endTime) - float64(startTime))/1000

	data := struct {
		Results     []map[string]string
		Query       string
		NextPageUrl string
		Info        map[string]interface{}
		FoundInfo   bool
		SearchTime float64
	}{
		Results:     results,
		Query:       query,
		NextPageUrl: next_page_url,
		Info:        wikiInfo,
		FoundInfo:   FoundInfo,
		SearchTime: totalSearchTimeSec,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
