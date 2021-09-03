package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// TestAdapterDisarm .
func TestAdapterDisarm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.DisarmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Disarm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterDisarmInternal(ctx, supportMock, actionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterDisarm .
func TestRequestErrorWhenAdapterDisarm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	actionMock := &actionServiceClientMock{}
	actionMock.On("Disarm", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterDisarmInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrRequest)
	a.Equal("disarm command error: request error", supportMock.message)
}

// TestResponseErrorWhenAdapterDisarm .
func TestResponseErrorWhenAdapterDisarm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.DisarmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Disarm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterDisarmInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrDisarmCommand)
	a.Equal("disarm command error: no disarm command success", supportMock.message)
}
