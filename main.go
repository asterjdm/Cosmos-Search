package main

import (
    _ "fmt"
    _"html"
    "log"
    "net/http"
	"example.com/Cosmos-Search/engines"
)

func main() {
	engines.search("haha")
    http.Handle("/", http.FileServer(http.Dir("./static/home")))

    log.Fatal(http.ListenAndServe(":5000", nil))

}

