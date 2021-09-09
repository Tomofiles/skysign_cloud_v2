package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "edge-px4/pkg/protos/action"
)

// TestAdapterTakeoff .
func TestAdapterTakeoff(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.TakeoffResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Takeoff", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterTakeoffInternal(ctx, actionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterTakeoff .
func TestRequestErrorWhenAdapterTakeoff(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	actionMock := &actionServiceClientMock{}
	actionMock.On("Takeoff", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterTakeoffInternal(ctx, actionMock)

	a.Equal("takeoff command error: request error", ret.Error())
}

// TestResponseErrorWhenAdapterTakeoff .
func TestResponseErrorWhenAdapterTakeoff(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.TakeoffResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Takeoff", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterTakeoffInternal(ctx, actionMock)

	a.Equal("takeoff command error: no takeoff command success", ret.Error())
}
