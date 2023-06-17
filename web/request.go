package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"errors"
	"bytes"
	"encoding/json"
)

const USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

func GetHtml(url string) (*bytes.Reader, error) {
	var client = &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", USER_AGENT)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := errors.New(fmt.Sprintf("failed to fetch data: %d %s", resp.StatusCode, resp.Status))
        return nil, err
    }

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(html)
	return reader, nil
}


func GetJson(url string){
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
	req.Header.Set("User-Agent", USER_AGENT)
	res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Print(err.Error())
    }

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
    if err != nil {
       	return nil, err
    }

	var data Data

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}