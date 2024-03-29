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

// TestPullUploadMission .
func TestPullUploadMission(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	respJson := json.Marshal(&skysign_proto.PullUploadMissionResponse{
		Id:        DefaultEdgeCommandID,
		CommandId: DefaultEdgeCommandID,
		MissionId: DefaultEdgeMissionID,
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resMethod = r.Method
		resPath = r.URL.Path
		resBody, _ = ioutil.ReadAll(r.Body)

		fmt.Fprintln(w, string(respJson))
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	upload, err := PullUploadMission(ts.URL, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/uploadmissions/command-id"
	expectBody := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("SEND   , Upload   , Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)
	expectMessage2 := fmt.Sprintf("RECEIVE, Upload   , data=%s\n", respJson)

	expectUpload := &model.UploadMission{
		ID:        DefaultEdgeCommandID,
		MissionID: DefaultEdgeMissionID,
	}

	a.Equal(expectMethod, resMethod)
	a.Equal(expectUrl, resPath)
	a.Equal(expectBody, resBody)

	a.Equal(expectUpload, upload)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestHttpClientErrorWhenPullUploadMission .
func TestHttpClientErrorWhenPullUploadMission(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	upload, err := PullUploadMission("$", support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/uploadmissions/command-id"
	expectBody := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	expectMessage := fmt.Sprintf("SEND   , Upload   , Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud upload http client error: http client do error: Post $/api/v1/communications/vehicle-id/uploadmissions/command-id: unsupported protocol scheme \"\""

	a.Nil(upload)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
}

// TestResponseJsonParseErrorWhenPullUploadMission .
func TestResponseJsonParseErrorWhenPullUploadMission(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	upload, err := PullUploadMission(ts.URL, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/uploadmissions/command-id"
	expectBody := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	expectMessage := fmt.Sprintf("SEND   , Upload   , Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud upload response error: unexpected EOF"

	a.Nil(upload)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
}
