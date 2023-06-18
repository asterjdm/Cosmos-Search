package main

import (
	"github.com/asterjdm/Cosmos-Search/web"
	"fmt"	
)


func GetWiki(query string) (map[string]interface{}, error) {
	var searchData []interface{}
	var encodedQuery string = web.UrlEncode(query)
	var searchUrl string = fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&search=%s&limit=1&namespace=0&format=json", encodedQuery)
	err := web.GetJson(searchUrl, &searchData)
	if err != nil {
		return nil, err
	}

	var string resultTitle = searchData[3][1]


	var string encodedResultTitle =  web.UrlEncode(resultTitle)

	var pageInfoUrl string = fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles=%s", encodedResultTitle)
	var pageData []interface{}

	err := web.GetJson(pageInfoUrl, &pageData)
	if err != nil {
		return nil, err
	}


	return nil, nil
}

func main() {
	GetWiki("chat")
}