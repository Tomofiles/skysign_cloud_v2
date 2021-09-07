package cloudlink

import (
	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/adapters/json"
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

	expectBody := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("Receive CLOUD data=%s\n", respJson)

	expectUpload := &edge.UploadMission{
		ID:        DefaultEdgeCommandID,
		MissionID: DefaultEdgeMissionID,
	}

	a.Equal(http.MethodPost, resMethod)
	a.Equal("/api/v1/communications/vehicle-id/uploadmissions/command-id", resPath)
	a.Equal(expectBody, resBody)

	a.Equal(expectUpload, upload)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestHttpClientErrorWhenPullUploadMission .
func TestHttpClientErrorWhenPullUploadMission(t *testing.T) {
	a := assert.New(t)

	dummyHost := "dummy-address.com"

	support := &supportMock{}

	upload, err := PullUploadMission("http://"+dummyHost, support, DefaultEdgeVehicleID, DefaultEdgeCommandID)

	expectBody := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("http client do error: Post http://%s/api/v1/communications/vehicle-id/uploadmissions/command-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)
	expectMessage3 := fmt.Sprintf("cloud upload http client error: Post http://%s/api/v1/communications/vehicle-id/uploadmissions/command-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)

	a.Nil(upload)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2, expectMessage3}, support.messages)
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

	expectBody := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := "cloud upload response error: unexpected EOF"

	a.Nil(upload)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}
