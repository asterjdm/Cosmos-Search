package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func search(query string) {
	var url string = "https://duckduckgo.com/search?q=" + query
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
        log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
    }

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
        log.Fatal(err)
    }

	title := doc.Find("title").Text()
    fmt.Println(title)
}


func main() {
	search("ss")
}