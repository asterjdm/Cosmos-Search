package main

import (
	"fmt"

	"github.com/asterjdm/Cosmos-Search/web"
)

type PageData struct {
	PageID  int    `json:"pageid"`
	NS      int    `json:"ns"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

func getFirstKey(m map[string]interface{}) (string, bool) {
	for key := range m {
		return key, true
	}
	return "", false
}

func GetWiki(query string) (*PageData, error) {
	var searchData []interface{}
	encodedQuery := web.UrlEncode(query)
	searchUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&search=%s&limit=1&namespace=0&format=json", encodedQuery)
	err := web.GetJson(searchUrl, &searchData)
	if err != nil {
		return nil, err
	}

	resultTitle := searchData[1].([]interface{})[0].(string)
	encodedResultTitle := web.UrlEncode(resultTitle)

	pageInfoUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles=%s", encodedResultTitle)

	var pageData struct {
		Query struct {
			Pages map[string]PageData `json:"pages"`
		} `json:"query"`
	}

	err = web.GetJson(pageInfoUrl, &pageData)
	if err != nil {
		return nil, err
	}

	wikiPageInfoPage := pageData.Query.Pages

	var firstPageKey string
	for k := range wikiPageInfoPage {
		firstPageKey = k
		break
	}


	wikiSummary := wikiPageInfoPage[firstPageKey].Extract
	wikiTitle := wikiPageInfoPage[firstPageKey].Title
	fmt.Println(wikiSummary)

	return nil, nil
}

func main() {
	_, err := GetWiki("underscore")
	if err != nil {
		fmt.Println(err)
	}
}
