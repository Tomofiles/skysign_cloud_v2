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

// TestGetUploadMission .
func TestGetUploadMission(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	respJson := json.Marshal(&skysign_proto.UploadMission{
		Id: DefaultEdgeMissionID,
		Waypoints: []*skysign_proto.Waypoint{
			{
				Latitude:       11.0,
				Longitude:      21.0,
				RelativeHeight: 31.0,
				Speed:          41.0,
			},
			{
				Latitude:       12.0,
				Longitude:      22.0,
				RelativeHeight: 32.0,
				Speed:          42.0,
			},
			{
				Latitude:       13.0,
				Longitude:      23.0,
				RelativeHeight: 33.0,
				Speed:          43.0,
			},
		},
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resMethod = r.Method
		resPath = r.URL.Path
		resBody, _ = ioutil.ReadAll(r.Body)

		fmt.Fprintln(w, string(respJson))
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	mission, err := GetUploadMission(ts.URL, support, DefaultEdgeMissionID)

	expectBody := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("Receive CLOUD data=%s\n", respJson)

	expectMission := &model.Mission{
		ID: DefaultEdgeMissionID,
		Waypoints: []*model.Waypoints{
			{
				Latitude:       11.0,
				Longitude:      21.0,
				RelativeHeight: 31.0,
				Speed:          41.0,
			},
			{
				Latitude:       12.0,
				Longitude:      22.0,
				RelativeHeight: 32.0,
				Speed:          42.0,
			},
			{
				Latitude:       13.0,
				Longitude:      23.0,
				RelativeHeight: 33.0,
				Speed:          43.0,
			},
		},
	}

	a.Equal(http.MethodGet, resMethod)
	a.Equal("/api/v1/uploadmissions/mission-id", resPath)
	a.Equal(expectBody, resBody)

	a.Equal(expectMission, mission)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestHttpClientErrorWhenGetUploadMission .
func TestHttpClientErrorWhenGetUploadMission(t *testing.T) {
	a := assert.New(t)

	dummyHost := "dummy-address.com"

	support := &supportMock{}

	mission, err := GetUploadMission("http://"+dummyHost, support, DefaultEdgeMissionID)

	expectBody := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("http client do error: Get http://%s/api/v1/uploadmissions/mission-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)
	expectMessage3 := fmt.Sprintf("cloud mission http client error: Get http://%s/api/v1/uploadmissions/mission-id: dial tcp: lookup %s: no such host", dummyHost, dummyHost)

	a.Nil(mission)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2, expectMessage3}, support.messages)
}

// TestResponseJsonParseErrorWhenGetUploadMission .
func TestResponseJsonParseErrorWhenGetUploadMission(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	mission, err := GetUploadMission(ts.URL, support, DefaultEdgeMissionID)

	expectBody := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := "cloud mission response error: unexpected EOF"

	a.Nil(mission)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}
