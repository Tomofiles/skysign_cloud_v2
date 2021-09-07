package cloudlink

import (
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/stretchr/testify/assert"
)

// TestPullCommand .
func TestPullCommand(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	respJson := json.Marshal(&skysign_proto.PullCommandResponse{
		Id:        DefaultEdgeVehicleID,
		CommandId: DefaultEdgeCommandID,
		Type:      skysign_proto.CommandType_ARM,
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

	expectBody := json.Marshal(&skysign_proto.PullCommandRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("Receive CLOUD data=%s\n", respJson)

	expectCommand := &model.Command{
		Type: "ARM",
	}

	a.Equal(http.MethodPost, resMethod)
	a.Equal("/api/v1/communications/vehicle-id/commands/command-id", resPath)
	a.Equal(expectBody, resBody)

	a.Equal(expectCommand, command)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestHttpClientErrorWhenPullCommand .
func TestHttpClientErrorWhenPullCommand(t *testing.T) {
	a := assert.New(t)

	dummyHost := "dummy-address.com"

	support := &supportMock{}

	command, err := PullCommand("http://"+dummyHost, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectBody := json.Marshal(&skysign_proto.PullCommandRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("http client do error: Post http://%s/api/v1/communications/vehicle-id/commands/command-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)
	expectMessage3 := fmt.Sprintf("cloud command http client error: Post http://%s/api/v1/communications/vehicle-id/commands/command-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)

	a.Nil(command)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2, expectMessage3}, support.messages)
}

// TestResponseJsonParseErrorWhenPullCommand .
func TestResponseJsonParseErrorWhenPullCommand(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	command, err := PullCommand(ts.URL, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectBody := json.Marshal(&skysign_proto.PullCommandRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := "cloud command response error: unexpected EOF"

	a.Nil(command)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}
