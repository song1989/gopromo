package utils

import (
	"io/ioutil"
	"net/http"
)

type Http struct {
}

func (this *Http) GetContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
