package main

import (
    _ "fmt"
    _"html"
    "log"
    "net/http"
)

func main() {

    http.Handle("/", http.FileServer(http.Dir("./static/home")))

    log.Fatal(http.ListenAndServe(":5000", nil))

}

