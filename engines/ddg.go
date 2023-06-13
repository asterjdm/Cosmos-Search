package main

import(
	"fmt"
	_ "net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"example.com/Cosmos-Search/web"
)

func Search(query string) {
	var url string = "https://duckduckgo.com?q=" + query
	/*var results = []struct{
		title string
		link string
		description string
	}{}*/
	html, err := web.GetHtml(url)
	if err != nil {
		//return results, err
	}
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
        log.Fatal(err)
    }

	title := doc.Find("title").Text()
    fmt.Println(title)
}

func main() {
	Search("hey")
}