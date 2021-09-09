package cloudlink

import (
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/model"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
)

// TestNoCommandIDsResponsePushTelemetry .
func TestNoCommandIDsResponsePushTelemetry(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	respJson := json.Marshal(&skysign_proto.PushTelemetryResponse{
		Id:         DefaultEdgeVehicleID,
		CommandIds: []string{},
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resMethod = r.Method
		resPath = r.URL.Path
		resBody, _ = ioutil.ReadAll(r.Body)

		fmt.Fprintln(w, string(respJson))
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	telemetry := model.NewTelemetry()
	telemetry.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&model.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&model.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&model.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&model.Armed{Armed: true})
	telemetry.SetFlightMode(&model.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry(ts.URL, support, telemetry)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/telemetry"
	expectBody := json.Marshal(&skysign_proto.PushTelemetryRequest{
		Id: DefaultEdgeVehicleID,
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            math.Sqrt(1.0*1.0 + 2.0*2.0),
			Armed:            true,
			FlightMode:       "XXX",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	})

	expectMessage1 := fmt.Sprintf("SEND   , Telemetry, Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)
	expectMessage2 := fmt.Sprintf("RECEIVE, Telemetry, data=%s\n", respJson)

	a.Equal(expectMethod, resMethod)
	a.Equal(expectUrl, resPath)
	a.Equal(expectBody, resBody)

	a.Equal(DefaultEdgeVehicleID, id)
	a.Empty(commandIDs.CommandIds)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestMultipleCommandIDsResponsePushTelemetry .
func TestMultipleCommandIDsResponsePushTelemetry(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultEdgeCommandID1 = DefaultEdgeCommandID + "-1"
		DefaultEdgeCommandID2 = DefaultEdgeCommandID + "-2"
		DefaultEdgeCommandID3 = DefaultEdgeCommandID + "-3"
	)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	respJson := json.Marshal(&skysign_proto.PushTelemetryResponse{
		Id:         DefaultEdgeVehicleID,
		CommandIds: []string{DefaultEdgeCommandID1, DefaultEdgeCommandID2, DefaultEdgeCommandID3},
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resMethod = r.Method
		resPath = r.URL.Path
		resBody, _ = ioutil.ReadAll(r.Body)

		fmt.Fprintln(w, string(respJson))
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	telemetry := model.NewTelemetry()
	telemetry.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&model.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&model.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&model.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&model.Armed{Armed: true})
	telemetry.SetFlightMode(&model.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry(ts.URL, support, telemetry)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/telemetry"
	expectBody := json.Marshal(&skysign_proto.PushTelemetryRequest{
		Id: DefaultEdgeVehicleID,
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            math.Sqrt(1.0*1.0 + 2.0*2.0),
			Armed:            true,
			FlightMode:       "XXX",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	})

	expectMessage1 := fmt.Sprintf("SEND   , Telemetry, Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)
	expectMessage2 := fmt.Sprintf("RECEIVE, Telemetry, data=%s\n", respJson)

	a.Equal(expectMethod, resMethod)
	a.Equal(expectUrl, resPath)
	a.Equal(expectBody, resBody)

	a.Equal(DefaultEdgeVehicleID, id)
	a.Equal([]string{DefaultEdgeCommandID1, DefaultEdgeCommandID2, DefaultEdgeCommandID3}, commandIDs.CommandIds)
	a.Nil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}

// TestNotPreparedWhenPushTelemetry .
func TestNotPreparedWhenPushTelemetry(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	respJson := json.Marshal(&skysign_proto.PushTelemetryResponse{
		Id:         DefaultEdgeVehicleID,
		CommandIds: []string{},
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(respJson))
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	tlm := model.NewTelemetry()
	tlm.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})

	id, commandIDs, err := PushTelemetry(ts.URL, support, tlm)

	expectError := "cloud telemetry request error: no telemetry prepared"

	a.Empty(id)
	a.Nil(commandIDs)
	a.Equal(expectError, err.Error())
	a.Empty(support.messages)
}

// TestHttpClientErrorWhenPushTelemetry .
func TestHttpClientErrorWhenPushTelemetry(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	telemetry := model.NewTelemetry()
	telemetry.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&model.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&model.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&model.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&model.Armed{Armed: true})
	telemetry.SetFlightMode(&model.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry("$", support, telemetry)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/telemetry"
	expectBody := json.Marshal(&skysign_proto.PushTelemetryRequest{
		Id: DefaultEdgeVehicleID,
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            math.Sqrt(1.0*1.0 + 2.0*2.0),
			Armed:            true,
			FlightMode:       "XXX",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	})

	expectMessage := fmt.Sprintf("SEND   , Telemetry, Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud telemetry http client error: http client do error: Post $/api/v1/communications/vehicle-id/telemetry: unsupported protocol scheme \"\""

	a.Empty(id)
	a.Empty(commandIDs)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
}

// TestResponseJsonParseErrorWhenPushTelemetry .
func TestResponseJsonParseErrorWhenPushTelemetry(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	telemetry := model.NewTelemetry()
	telemetry.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&model.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&model.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&model.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&model.Armed{Armed: true})
	telemetry.SetFlightMode(&model.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry(ts.URL, support, telemetry)

	expectMethod := http.MethodPost
	expectUrl := "/api/v1/communications/vehicle-id/telemetry"
	expectBody := json.Marshal(&skysign_proto.PushTelemetryRequest{
		Id: DefaultEdgeVehicleID,
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            math.Sqrt(1.0*1.0 + 2.0*2.0),
			Armed:            true,
			FlightMode:       "XXX",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	})

	expectMessage := fmt.Sprintf("SEND   , Telemetry, Method=%s, API=%s, Message=%s", expectMethod, expectUrl, expectBody)

	expectError := "cloud telemetry response error: unexpected EOF"

	a.Empty(id)
	a.Empty(commandIDs)
	a.Equal(expectError, err.Error())
	a.Equal([]string{expectMessage}, support.messages)
}
