package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpClientDo(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resMethod = r.Method
		resPath = r.URL.Path
		resBody, _ = ioutil.ReadAll(r.Body)

		fmt.Fprintln(w, "{}")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	respBody, err := HttpClientDo(support, http.MethodPost, ts.URL+"/path/aaa", []byte("{}"))

	a.Equal(http.MethodPost, resMethod)
	a.Equal("/path/aaa", resPath)
	a.Equal([]byte("{}"), resBody)

	a.Equal([]byte("{}\n"), respBody)
	a.Nil(err)
	a.Empty(support.messages)
}

func TestCLientDoErrorWhenHttpClientDo(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{}")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	respBody, err := HttpClientDo(support, http.MethodGet, "", []byte("{}"))

	expectMessage := "http client do error: Get : unsupported protocol scheme \"\""

	a.Nil(respBody)
	a.NotNil(err)
	a.Equal([]string{expectMessage}, support.messages)
}
