package engines

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"example.com/Cosmos-Search/web"
)

func Search(query string, page *int) ([]map[string]string, error) {
	var pageNumber int = 0
	if page != nil {
		pageNumber = *page
	}

	encodedQuery := web.UrlEncode(query)
	url := fmt.Sprintf("https://google.com/search?q=%s&start=%d", encodedQuery, pageNumber*10)
	results := []map[string]string{}

	html, err := web.GetHtml(url)
	if err != nil {
		return results, err
	}

	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return results, err
	}

	doc.Find(".VwiC3b").Each(func(i int, s *goquery.Selection) {
		title := s.Parents().First().Parents().First().Find("h3").Text()
		link, _ := s.Parents().First().Parents().First().Find("a").Attr("href")
		description := s.Text()

		result := map[string]string{
			"title":       title,
			"link":        link,
			"description": description,
		}
		results = append(results, result)
	})

	return results, nil
}
