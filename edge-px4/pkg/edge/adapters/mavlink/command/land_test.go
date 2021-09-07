package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "edge-px4/pkg/protos/action"
)

// TestAdapterLand .
func TestAdapterLand(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.LandResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Land", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterLandInternal(ctx, supportMock, actionMock)

	a.Nil(ret)
	a.Empty(supportMock.message)
}

// TestRequestErrorWhenAdapterLand .
func TestRequestErrorWhenAdapterLand(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	actionMock := &actionServiceClientMock{}
	actionMock.On("Land", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterLandInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrRequest)
	a.Equal("land command error: request error", supportMock.message)
}

// TestResponseErrorWhenAdapterLand .
func TestResponseErrorWhenAdapterLand(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.LandResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Land", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterLandInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrLandCommand)
	a.Equal("land command error: no land command success", supportMock.message)
}
