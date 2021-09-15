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
				Latitude:         11.0,
				Longitude:        21.0,
				RelativeAltitude: 31.0,
				Speed:            41.0,
			},
			{
				Latitude:         12.0,
				Longitude:        22.0,
				RelativeAltitude: 32.0,
				Speed:            42.0,
			},
			{
				Latitude:         13.0,
				Longitude:        23.0,
				RelativeAltitude: 33.0,
				Speed:            43.0,
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

	expectMethod := http.MethodGet
	expectUrl := "/api/v1/uploadmissions/mission-id"
	expectBody := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	expectMessage1 := fmt.Sprintf("SEND   , Mission  , Method=%s , API=%s, Message=%s", expectMethod, expectUrl, expectBody)
	expectMessage2 := fmt.Sprintf("RECEIVE, Mission  , data=%s\n", respJson)

	expectMission := &model.Mission{
		ID: DefaultEdgeMissionID,
		Waypoints: []*model.Waypoints{
			{
				LatitudeDegree:    11.0,
				LongitudeDegree:   21.0,
				RelativeAltitudeM: 31.0,
				SpeedMS:           41.0,
			},
			{
				LatitudeDegree:    12.0,
				LongitudeDegree:   22.0,
				RelativeAltitudeM: 32.0,
				SpeedMS:           42.0,
			},
			{
				LatitudeDegree:    13.0,
				LongitudeDegree:   23.0,
				RelativeAltitudeM: 33.0,
				SpeedMS:           43.0,
			},
		},
	}

	a.Equal(expectMethod, resMethod)
	a.Equal(expectUrl, resPath)
	a.Equal(expectBody, resBody)

	a.Equal(expectMission, mission)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestHttpClientErrorWhenGetUploadMission .
func TestHttpClientErrorWhenGetUploadMission(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	mission, err := GetUploadMission("$", support, DefaultEdgeMissionID)

	expectMethod := http.MethodGet
	expectUrl := "/api/v1/uploadmissions/mission-id"
	expectBody := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	expectMessage := fmt.Sprintf("SEND   , Mission  , Method=%s , API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud mission http client error: http client do error: Get $/api/v1/uploadmissions/mission-id: unsupported protocol scheme \"\""

	a.Nil(mission)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
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

	expectMethod := http.MethodGet
	expectUrl := "/api/v1/uploadmissions/mission-id"
	expectBody := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	expectMessage := fmt.Sprintf("SEND   , Mission  , Method=%s , API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud mission response error: unexpected EOF"

	a.Nil(mission)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
}
