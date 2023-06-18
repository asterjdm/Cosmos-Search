package main

import (
	"github.com/asterjdm/Cosmos-Search/web"
	"fmt"	
)


func GetWiki(query string) (map[string]interface{}, error) {
	var data []interface{}
	var url string = fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&search=%s&limit=1&namespace=0&format=json", query)
	err := web.GetJson(url, &data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data[3])


	return nil, nil
}

func main() {
	GetWiki("chat")
}