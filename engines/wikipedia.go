package engines

import (
	"github.com/trietmn/go-wiki"
)

func getWiki(query string) (map[string]interface{}, error) {
	searchResult, _, err := gowiki.Search(query, 1, true)
	if err != nil {
		return nil, err
	}

	wikiInfo := make(map[string]interface{})

	if len(searchResult) <= 0 {
		wikiInfo["Found"] = false
		wikiInfo["Summary"] = ""
		return wikiInfo, nil
	}

	page, err := gowiki.GetPage(searchResult[0], -1, false, true)
	if err != nil {
		return nil, err
	}

	summary, err := gowiki.Summary(searchResult[0], 3, -1, false, true)
	if err != nil {
		return nil, err
	}

	wikiInfo["Found"] = true
	wikiInfo["Summary"] = summary

	return wikiInfo, nil
}
