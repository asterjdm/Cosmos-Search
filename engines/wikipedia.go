package main

import (
	"github.com/asterjdm/Cosmos-Search/web"
	"fmt"	
)

type Pages struct {
	PageID int `json:"pageid"`
	NS     int `json:"ns"`
	Title  string `json:"title"`
	Extract string `json:"extract"`
}


func GetWiki(query string) (map[string]interface{}, error) {
	var searchData []interface{}
	var encodedQuery string = web.UrlEncode(query)
	var searchUrl string = fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&search=%s&limit=1&namespace=0&format=json", encodedQuery)
	err := web.GetJson(searchUrl, &searchData)
	if err != nil {
		return nil, err
	}

	var resultTitle string = searchData[1].([]interface{})[0].(string)


	var encodedResultTitle string =  web.UrlEncode(resultTitle)

	var pageInfoUrl string = fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles=%s", encodedResultTitle)
	var pageData struct

	err = web.GetJson(pageInfoUrl, &pageData)
	if err != nil {
		return nil, err
	}

	fmt.Println(pageData)


	return nil, nil
}

func main() {
	_, err := GetWiki("chat")
	if err != nil {
		fmt.Println(err)
	}
}