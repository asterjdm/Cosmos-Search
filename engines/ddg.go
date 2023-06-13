package main

import(
	"fmt"
	_ "net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"example.com/Cosmos-Search/web"
)

func Search(query string) ([]struct {
	title       string
	link        string
	description string
}, error) {

	var url string = "https://google.com/search?q=" + query
	var results = []struct{
		title string
		link string
		description string
	}{}


	html, err := web.GetHtml(url)
	if err != nil {
		return results, err
	}
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
        return results, err
    }

	doc.Find("#search").Each(func(i int, s *goquery.Selection) {

	})
    //fmt.Println(title)
}

func main() {
	Search("github")
}