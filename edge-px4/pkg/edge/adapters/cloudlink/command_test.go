package cloudlink

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/json"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"

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

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/commands/command-id"
	expectBody := json.Marshal(&skysign_proto.PullCommandRequest{})

	expectMessage1 := fmt.Sprintf("SEND   , Command  , Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)
	expectMessage2 := fmt.Sprintf("RECEIVE, Command  , data=%s\n", respJson)

	expectCommand := &model.Command{
		Type: "ARM",
	}

	a.Equal(expectMethod, resMethod)
	a.Equal(expectUrl, resPath)
	a.Equal(expectBody, resBody)

	a.Equal(expectCommand, command)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestHttpClientErrorWhenPullCommand .
func TestHttpClientErrorWhenPullCommand(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	command, err := PullCommand("$", support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/commands/command-id"
	expectBody := json.Marshal(&skysign_proto.PullCommandRequest{})

	expectMessage := fmt.Sprintf("SEND   , Command  , Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud command http client error: http client do error: Post $/api/v1/communications/vehicle-id/commands/command-id: unsupported protocol scheme \"\""

	a.Nil(command)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
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

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/commands/command-id"
	expectBody := json.Marshal(&skysign_proto.PullCommandRequest{})

	expectMessage := fmt.Sprintf("SEND   , Command  , Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud command response error: unexpected EOF"

	a.Nil(command)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
}
