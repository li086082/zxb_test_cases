package util

import (
	"io"
	"net/http"
	"strings"
)

func Request(method string, url string, header map[string]string, body string) (string, error) {
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	for key, value := range header {
		request.Header.Add(key, value)
	}
	client := &http.Client{}
	response, err2 := client.Do(request)
	if err2 != nil {
		return "", err2
	}
	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			return
		}
	}(response.Body)

	all, err4 := io.ReadAll(response.Body)
	if err4 != nil {
		return "", err4
	}
	return string(all), nil
}
