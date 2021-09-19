package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/action"
)

// TestAdapterDisarm .
func TestAdapterDisarm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.DisarmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Disarm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterDisarmInternal(ctx, actionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterDisarm .
func TestRequestErrorWhenAdapterDisarm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	actionMock := &actionServiceClientMock{}
	actionMock.On("Disarm", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterDisarmInternal(ctx, actionMock)

	a.Equal("disarm command error: request error", ret.Error())
}

// TestResponseErrorWhenAdapterDisarm .
func TestResponseErrorWhenAdapterDisarm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.DisarmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Disarm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterDisarmInternal(ctx, actionMock)

	a.Equal("disarm command error: no disarm command success", ret.Error())
}
