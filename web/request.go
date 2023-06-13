package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

func getHtml(url string) []byte {
	var client = &http.Client{}
	const USER_AGENT = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", USER_AGENT)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Sprintf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
        return nil, err
    }

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", html)
	return html
}

