package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Res struct{}

func main() {
	result := &Res{}
	headers := &http.Header{}
	GetJsonWithHeader("https://api.bilibili.com/x/web-interface/nav", headers, result)
}

func GetJsonWithHeader(url string, headers *http.Header, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header = *headers
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println(result)
		_ = resp.Body.Close()
	}()
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}
	return nil
}
