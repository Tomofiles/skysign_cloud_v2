package http

import (
	"edge-px4/pkg/edge/domain/common"
	"io/ioutil"
	"net/http"
	"strings"
)

// HttpClientDo .
func HttpClientDo(
	support common.Support,
	method string,
	url string,
	reqBody []byte,
) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(string(reqBody)))
	if err != nil {
		support.NotifyError("http request error: %v", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		support.NotifyError("http client do error: %v", err)
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		support.NotifyError("http response error: %v", err)
		return nil, err
	}

	return respBody, nil
}
