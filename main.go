package main

import (
    "fmt"
    _"html"
    "log"
    "net/http"
	"example.com/Cosmos-Search/engines"
)

func main() {
	results, err := engines.Search("golang programming language")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
    http.Handle("/", http.FileServer(http.Dir("./static/home")))

    log.Fatal(http.ListenAndServe(":5000", nil))

}

