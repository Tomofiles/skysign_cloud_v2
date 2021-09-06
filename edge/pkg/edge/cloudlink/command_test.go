package cloudlink

import (
	"edge/pkg/edge"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullCommand(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	respJson, _ := json.Marshal(edge.Command{
		Type: "XXX",
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resMethod = r.Method
		resPath = r.URL.Path
		resBody, _ = ioutil.ReadAll(r.Body)

		fmt.Fprintln(w, string(respJson))
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	command, err := PullCommand(ts.URL, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", "{}")
	expectMessage2 := fmt.Sprintf("Receive CLOUD data=%s\n", respJson)

	expectCommand := &edge.Command{
		Type: "XXX",
	}

	a.Equal(http.MethodPost, resMethod)
	a.Equal("/api/v1/communications/vehicle-id/commands/command-id", resPath)
	a.Equal([]byte("{}"), resBody)

	a.Equal(expectCommand, command)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

func TestHttpClientErrorWhenPullCommand(t *testing.T) {
	a := assert.New(t)

	dummyHost := "dummy-address.com"

	support := &supportMock{}

	command, err := PullCommand("http://"+dummyHost, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", "{}")
	expectMessage2 := fmt.Sprintf("http client do error: Post http://%s/api/v1/communications/vehicle-id/commands/command-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)
	expectMessage3 := fmt.Sprintf("cloud command http client error: Post http://%s/api/v1/communications/vehicle-id/commands/command-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)

	a.Nil(command)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2, expectMessage3}, support.messages)
}

func TestResponseJsonParseErrorWhenPullCommand(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	respJson, _ := json.Marshal(edge.Command{
		Type: "XXX",
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(respJson)+"{")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	command, err := PullCommand(ts.URL, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", "{}")
	expectMessage2 := "cloud command response error: invalid character '{' after top-level value"

	a.Nil(command)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}
