package proto

import (
	"errors"
	"strings"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/stretchr/testify/assert"
)

// TestValidateCreateVehicleRequest_Success .
func TestValidateCreateVehicleRequest_Success(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 200)
	communicationID := strings.Repeat("X", 36)

	req := &skysign_proto.Vehicle{
		Name:            name,
		CommunicationId: communicationID,
	}
	ret := ValidateCreateVehicleRequest(req)

	a.Nil(ret)
}

// TestValidateCreateVehicleRequest_Failure_Blank .
func TestValidateCreateVehicleRequest_Failure_Blank(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 0)
	communicationID := strings.Repeat("X", 0)

	req := &skysign_proto.Vehicle{
		Name:            name,
		CommunicationId: communicationID,
	}
	ret := ValidateCreateVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 2)
	a.Equal("cannot be blank", errs["name"].Error())
	a.Equal("cannot be blank", errs["communication_id"].Error())
}

// TestValidateCreateVehicleRequest_Failure_Length .
func TestValidateCreateVehicleRequest_Failure_Length(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 201)
	communicationID := strings.Repeat("X", 37)

	req := &skysign_proto.Vehicle{
		Name:            name,
		CommunicationId: communicationID,
	}
	ret := ValidateCreateVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 2)
	a.Equal("the length must be no more than 200", errs["name"].Error())
	a.Equal("the length must be no more than 36", errs["communication_id"].Error())
}
