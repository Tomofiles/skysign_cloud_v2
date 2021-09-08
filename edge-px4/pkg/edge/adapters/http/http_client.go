package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HttpClientDo .
func HttpClientDo(
	method string,
	url string,
	reqBody []byte,
) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, fmt.Errorf("http request error: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client do error: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("http response error: %w", err)
	}

	return respBody, nil
}
