package main

import(
	"github.com/trietmn/go-wiki"
	"fmt"
	"strings"
)

func getWiki(query string) (map[string]string, error) {
	search_result, _, err := gowiki.Search(query, 1, true)
	if err != nil {
		return nil, err
	}
	if len(search_result) <= 0 {
		return nil, nil
	}

	page, err := gowiki.GetPage(search_result[0], -1, false, true)
	if err != nil {
		return nil, err
	}
	summary, err := gowiki.Summary(search_result[0], 3, -1, false, true)
	if err != nil {
		return nil, err
	}
	image, err := page.GetImagesURL()
	fmt.Println(image)
	wikiInfo := map[string]string{
		"Summary": summary,
		"Image": image[0],
	}
	return wikiInfo, nil
}

func main() {
	wikiInfo, _ := getWiki("google")
	fmt.Println(wikiInfo["Summary"])
	fmt.Println(wikiInfo["Image"])
}