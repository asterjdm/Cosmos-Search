package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"example.com/Cosmos-Search/web"
)

func Search(query string) ([]map[string]string, error) {

	var url string = "https://google.com/search?q=" + query
	var results = []map[string]string{}

	html, err := web.GetHtml(url)
	if err != nil {
		return results, err
	}
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return results, err
	}

	doc.Find("#search").
		Children().First().
		Children().Last().
		Children().Last().
		Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Attr("class"))
			title := s.Find("h3").Text()
			link, _ := s.Find("a").Attr("href")
			description := s.Find(".VwiC3b").Children().Last().Text()

			result := map[string]string{
				"title":       title,
				"link":        link,
				"description": description,
			}
			results = append(results, result)
		})

	return results, nil
}

func main() {
	results, err := Search("github")
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println("Title:", result["title"])
		fmt.Println("Link:", result["link"])
		fmt.Println("Description:", result["description"])
		fmt.Println()
	}
}
