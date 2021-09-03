package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// TestAdapterArm .
func TestAdapterArm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.ArmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Arm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterArmInternal(ctx, supportMock, actionMock)

	a.Nil(ret)
	a.Empty(supportMock.message)
}

// TestRequestErrorWhenAdapterArm .
func TestRequestErrorWhenAdapterArm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	actionMock := &actionServiceClientMock{}
	actionMock.On("Arm", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterArmInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrRequest)
	a.Equal("arm command error: request error", supportMock.message)
}

// TestResponseErrorWhenAdapterArm .
func TestResponseErrorWhenAdapterArm(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.ArmResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Arm", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterArmInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrArmCommand)
	a.Equal("arm command error: no arm command success", supportMock.message)
}
