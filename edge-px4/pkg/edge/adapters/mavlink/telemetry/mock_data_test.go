package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_core "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/core"
	mavsdk_rpc_telemetry "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/telemetry"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

var (
	ErrRequest = errors.New("request error")
	ErrReceive = errors.New("receive error")
	ErrClose   = errors.New("close error")
)

type supportMock struct {
	messages []string
}

func (m *supportMock) NotifyInfo(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

func (m *supportMock) NotifyError(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

type coreServiceClientMock struct {
	mavsdk_rpc_core.CoreServiceClient
	mock.Mock
}

func (m *coreServiceClientMock) SubscribeConnectionState(ctx context.Context, in *mavsdk_rpc_core.SubscribeConnectionStateRequest, opts ...grpc.CallOption) (mavsdk_rpc_core.CoreService_SubscribeConnectionStateClient, error) {
	ret := m.Called(in)
	var v mavsdk_rpc_core.CoreService_SubscribeConnectionStateClient
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(mavsdk_rpc_core.CoreService_SubscribeConnectionStateClient)
	}
	return v, ret.Error(1)
}

type telemetryServiceClientMock struct {
	mavsdk_rpc_telemetry.TelemetryServiceClient
	mock.Mock
}

func (m *telemetryServiceClientMock) SubscribeArmed(ctx context.Context, in *mavsdk_rpc_telemetry.SubscribeArmedRequest, opts ...grpc.CallOption) (mavsdk_rpc_telemetry.TelemetryService_SubscribeArmedClient, error) {
	ret := m.Called(in)
	var v mavsdk_rpc_telemetry.TelemetryService_SubscribeArmedClient
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(mavsdk_rpc_telemetry.TelemetryService_SubscribeArmedClient)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientMock) SubscribeFlightMode(ctx context.Context, in *mavsdk_rpc_telemetry.SubscribeFlightModeRequest, opts ...grpc.CallOption) (mavsdk_rpc_telemetry.TelemetryService_SubscribeFlightModeClient, error) {
	ret := m.Called(in)
	var v mavsdk_rpc_telemetry.TelemetryService_SubscribeFlightModeClient
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(mavsdk_rpc_telemetry.TelemetryService_SubscribeFlightModeClient)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientMock) SubscribePosition(ctx context.Context, in *mavsdk_rpc_telemetry.SubscribePositionRequest, opts ...grpc.CallOption) (mavsdk_rpc_telemetry.TelemetryService_SubscribePositionClient, error) {
	ret := m.Called(in)
	var v mavsdk_rpc_telemetry.TelemetryService_SubscribePositionClient
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(mavsdk_rpc_telemetry.TelemetryService_SubscribePositionClient)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientMock) SubscribeAttitudeQuaternion(ctx context.Context, in *mavsdk_rpc_telemetry.SubscribeAttitudeQuaternionRequest, opts ...grpc.CallOption) (mavsdk_rpc_telemetry.TelemetryService_SubscribeAttitudeQuaternionClient, error) {
	ret := m.Called(in)
	var v mavsdk_rpc_telemetry.TelemetryService_SubscribeAttitudeQuaternionClient
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(mavsdk_rpc_telemetry.TelemetryService_SubscribeAttitudeQuaternionClient)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientMock) SubscribeGroundSpeedNed(ctx context.Context, in *mavsdk_rpc_telemetry.SubscribeGroundSpeedNedRequest, opts ...grpc.CallOption) (mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient, error) {
	ret := m.Called(in)
	var v mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient)
	}
	return v, ret.Error(1)
}

type coreServiceClientConnectionStateMock struct {
	grpc.ClientStream
	mock.Mock
	i int
}

func (m *coreServiceClientConnectionStateMock) Recv() (*mavsdk_rpc_core.ConnectionStateResponse, error) {
	defer func() {
		m.i++
	}()
	ret := m.Called(m.i)
	var v *mavsdk_rpc_core.ConnectionStateResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_core.ConnectionStateResponse)
	}
	return v, ret.Error(1)
}

func (m *coreServiceClientConnectionStateMock) CloseSend() error {
	ret := m.Called()
	return ret.Error(0)
}

type telemetryServiceClientArmedMock struct {
	grpc.ClientStream
	mock.Mock
	i int
}

func (m *telemetryServiceClientArmedMock) Recv() (*mavsdk_rpc_telemetry.ArmedResponse, error) {
	defer func() {
		m.i++
	}()
	ret := m.Called(m.i)
	var v *mavsdk_rpc_telemetry.ArmedResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_telemetry.ArmedResponse)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientArmedMock) CloseSend() error {
	ret := m.Called()
	return ret.Error(0)
}

type telemetryServiceClientFlightModeMock struct {
	grpc.ClientStream
	mock.Mock
	i int
}

func (m *telemetryServiceClientFlightModeMock) Recv() (*mavsdk_rpc_telemetry.FlightModeResponse, error) {
	defer func() {
		m.i++
	}()
	ret := m.Called(m.i)
	var v *mavsdk_rpc_telemetry.FlightModeResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_telemetry.FlightModeResponse)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientFlightModeMock) CloseSend() error {
	ret := m.Called()
	return ret.Error(0)
}

type telemetryServiceClientPositionMock struct {
	grpc.ClientStream
	mock.Mock
	i int
}

func (m *telemetryServiceClientPositionMock) Recv() (*mavsdk_rpc_telemetry.PositionResponse, error) {
	defer func() {
		m.i++
	}()
	ret := m.Called(m.i)
	var v *mavsdk_rpc_telemetry.PositionResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_telemetry.PositionResponse)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientPositionMock) CloseSend() error {
	ret := m.Called()
	return ret.Error(0)
}

type telemetryServiceClientQuaternionMock struct {
	grpc.ClientStream
	mock.Mock
	i int
}

func (m *telemetryServiceClientQuaternionMock) Recv() (*mavsdk_rpc_telemetry.AttitudeQuaternionResponse, error) {
	defer func() {
		m.i++
	}()
	ret := m.Called(m.i)
	var v *mavsdk_rpc_telemetry.AttitudeQuaternionResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_telemetry.AttitudeQuaternionResponse)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientQuaternionMock) CloseSend() error {
	ret := m.Called()
	return ret.Error(0)
}

type telemetryServiceClientVelocityMock struct {
	grpc.ClientStream
	mock.Mock
	i int
}

func (m *telemetryServiceClientVelocityMock) Recv() (*mavsdk_rpc_telemetry.GroundSpeedNedResponse, error) {
	defer func() {
		m.i++
	}()
	ret := m.Called(m.i)
	var v *mavsdk_rpc_telemetry.GroundSpeedNedResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_telemetry.GroundSpeedNedResponse)
	}
	return v, ret.Error(1)
}

func (m *telemetryServiceClientVelocityMock) CloseSend() error {
	ret := m.Called()
	return ret.Error(0)
}
