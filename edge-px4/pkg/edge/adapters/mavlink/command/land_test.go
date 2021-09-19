package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/action"
)

// TestAdapterLand .
func TestAdapterLand(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.LandResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Land", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterLandInternal(ctx, actionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterLand .
func TestRequestErrorWhenAdapterLand(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	actionMock := &actionServiceClientMock{}
	actionMock.On("Land", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterLandInternal(ctx, actionMock)

	a.Equal("land command error: request error", ret.Error())
}

// TestResponseErrorWhenAdapterLand .
func TestResponseErrorWhenAdapterLand(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.LandResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Land", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterLandInternal(ctx, actionMock)

	a.Equal("land command error: no land command success", ret.Error())
}
