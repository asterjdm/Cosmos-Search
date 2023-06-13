package engines

import (
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

	doc.Find(".kvH3mc").Each(func(i int, s *goquery.Selection) {
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

