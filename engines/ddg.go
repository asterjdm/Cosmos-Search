package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func search(query string) {
	var url string = "https://duckduckgo.com/search?q=" + query
	var client = &http.Client{}
	const USER_AGENT = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", USER_AGENT)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)

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