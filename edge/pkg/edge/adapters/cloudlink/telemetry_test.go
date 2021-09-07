package cloudlink

import (
	"edge/pkg/edge"
	"edge/pkg/edge/domain/telemetry"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNoCommandIDsResponsePushTelemetry .
func TestNoCommandIDsResponsePushTelemetry(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	var resMethod, resPath string
	var resBody []byte
	respJson, _ := json.Marshal(edge.CommandIDs{
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

	telemetry := telemetry.NewTelemetry()
	telemetry.SetConnectionState(&edge.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&edge.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&edge.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&edge.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&edge.Armed{Armed: true})
	telemetry.SetFlightMode(&edge.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry(ts.URL, support, telemetry)

	expectBody, _ := json.Marshal(edge.Telemetry{
		ID: DefaultEdgeVehicleID,
		State: &edge.State{
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

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("Receive CLOUD data=%s\n", respJson)

	a.Equal(http.MethodPost, resMethod)
	a.Equal("/api/v1/communications/vehicle-id/telemetry", resPath)
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
	respJson, _ := json.Marshal(edge.CommandIDs{
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

	telemetry := telemetry.NewTelemetry()
	telemetry.SetConnectionState(&edge.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&edge.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&edge.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&edge.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&edge.Armed{Armed: true})
	telemetry.SetFlightMode(&edge.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry(ts.URL, support, telemetry)

	expectBody, _ := json.Marshal(edge.Telemetry{
		ID: DefaultEdgeVehicleID,
		State: &edge.State{
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

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("Receive CLOUD data=%s\n", respJson)

	a.Equal(http.MethodPost, resMethod)
	a.Equal("/api/v1/communications/vehicle-id/telemetry", resPath)
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

	respJson, _ := json.Marshal(edge.CommandIDs{
		CommandIds: []string{},
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(respJson))
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	tlm := telemetry.NewTelemetry()
	tlm.SetConnectionState(&edge.ConnectionState{VehicleID: DefaultEdgeVehicleID})

	id, commandIDs, err := PushTelemetry(ts.URL, support, tlm)

	expectMessage1 := "cloud telemetry request error: no telemetry prepared"

	a.Empty(id)
	a.Nil(commandIDs)
	a.Equal(telemetry.ErrNotPrepared, err)
	a.Equal([]string{expectMessage1}, support.messages)
}

// TestHttpClientErrorWhenPushTelemetry .
func TestHttpClientErrorWhenPushTelemetry(t *testing.T) {
	a := assert.New(t)

	dummyHost := "dummy-address.com"

	support := &supportMock{}

	telemetry := telemetry.NewTelemetry()
	telemetry.SetConnectionState(&edge.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&edge.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&edge.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&edge.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&edge.Armed{Armed: true})
	telemetry.SetFlightMode(&edge.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry("http://"+dummyHost, support, telemetry)

	expectBody, _ := json.Marshal(edge.Telemetry{
		ID: DefaultEdgeVehicleID,
		State: &edge.State{
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

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := fmt.Sprintf("http client do error: Post http://%s/api/v1/communications/vehicle-id/telemetry: dial tcp: lookup %s: no such host", dummyHost, dummyHost)
	expectMessage3 := fmt.Sprintf("cloud telemetry http client error: Post http://%s/api/v1/communications/vehicle-id/telemetry: dial tcp: lookup %s: no such host", dummyHost, dummyHost)

	a.Empty(id)
	a.Empty(commandIDs)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2, expectMessage3}, support.messages)
}

// TestResponseJsonParseErrorWhenPushTelemetry .
func TestResponseJsonParseErrorWhenPushTelemetry(t *testing.T) {
	a := assert.New(t)

	support := &supportMock{}

	respJson, _ := json.Marshal(edge.CommandIDs{
		CommandIds: []string{},
	})
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(respJson)+"{")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	telemetry := telemetry.NewTelemetry()
	telemetry.SetConnectionState(&edge.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	telemetry.SetPosition(&edge.Position{Latitude: 1.0, Longitude: 2.0, Altitude: 3.0, RelativeAltitude: 4.0})
	telemetry.SetQuaternion(&edge.Quaternion{X: 6.0, Y: 7.0, Z: 8.0, W: 9.0})
	telemetry.SetVelocity(&edge.Velocity{North: 1.0, East: 2.0, Down: 3.0})
	telemetry.SetArmed(&edge.Armed{Armed: true})
	telemetry.SetFlightMode(&edge.FlightMode{FlightMode: "XXX"})

	id, commandIDs, err := PushTelemetry(ts.URL, support, telemetry)

	expectBody, _ := json.Marshal(edge.Telemetry{
		ID: DefaultEdgeVehicleID,
		State: &edge.State{
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

	expectMessage1 := fmt.Sprintf("Send CLOUD data=%s", expectBody)
	expectMessage2 := "cloud telemetry response error: invalid character '{' after top-level value"

	a.Empty(id)
	a.Empty(commandIDs)
	a.NotNil(err)
	a.Equal([]string{expectMessage1, expectMessage2}, support.messages)
}
