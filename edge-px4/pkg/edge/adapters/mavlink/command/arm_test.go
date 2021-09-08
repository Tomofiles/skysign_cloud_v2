package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "edge-px4/pkg/protos/action"
)

// TestAdapterArm .
func TestAdapterArm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.ArmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Arm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterArmInternal(ctx, actionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterArm .
func TestRequestErrorWhenAdapterArm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	actionMock := &actionServiceClientMock{}
	actionMock.On("Arm", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterArmInternal(ctx, actionMock)

	a.Equal("arm command error: request error", ret.Error())
}

// TestResponseErrorWhenAdapterArm .
func TestResponseErrorWhenAdapterArm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.ArmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Arm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterArmInternal(ctx, actionMock)

	a.Equal("arm command error: no arm command success", ret.Error())
}
