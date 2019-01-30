package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(getUrl string, params url.Values, token string) ([]byte, int, error) {
	u, _ := url.Parse(getUrl)
	u.RawQuery = params.Encode()
	data, code, err := httpDo("GET", u.String(), []byte(""), token)
	if err != nil {
		return nil, code, err
	}
	return data, code, nil
}

func HttpPost(postUrl string, body []byte, token string) ([]byte, int, error) {
	u, _ := url.Parse(postUrl)
	data, code, err := httpDo("POST", u.String(), body, token)
	if err != nil {
		return nil, code, err
	}
	return data, code, nil
}

func httpDo(methodType string, url string, body []byte, token string) ([]byte, int, error) {
	fmt.Println(url)
	if methodType == "" {
		methodType = "GET"
	}

	client := &http.Client{}
	req, err := http.NewRequest(methodType, url, bytes.NewReader(body))
	if err != nil {
		return nil, 500, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return data, resp.StatusCode, nil
}
