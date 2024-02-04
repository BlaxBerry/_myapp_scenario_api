package tools

import (
	"io"
	"net/http"
)

func Request(url string) (error, *string, *int) {
	response, err := http.Get(url)
	if err != nil {
		return err, nil, &response.StatusCode
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err, nil, &response.StatusCode
	}

	stringBody := string(body)
	return nil, &stringBody, &response.StatusCode
}
