package proto

import (
	"errors"
	"strings"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestValidateGetVehicleRequest_Success .
func TestValidateGetVehicleRequest_Success(t *testing.T) {
	a := assert.New(t)

	id, _ := uuid.NewRandom()

	req := &skysign_proto.GetVehicleRequest{
		Id: id.String(),
	}
	ret := ValidateGetVehicleRequest(req)

	a.Nil(ret)
}

// TestValidateGetVehicleRequest_Failure_Blank .
func TestValidateGetVehicleRequest_Failure_Blank(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 0)

	req := &skysign_proto.GetVehicleRequest{
		Id: id,
	}
	ret := ValidateGetVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("cannot be blank", errs["id"].Error())
}

// TestValidateGetVehicleRequest_Failure_Length .
func TestValidateGetVehicleRequest_Failure_Length(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 37)

	req := &skysign_proto.GetVehicleRequest{
		Id: id,
	}
	ret := ValidateGetVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("the length must be exactly 36", errs["id"].Error())
}

// TestValidateGetVehicleRequest_Failure_IllegalFormat .
func TestValidateGetVehicleRequest_Failure_IllegalFormat(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 36)

	req := &skysign_proto.GetVehicleRequest{
		Id: id,
	}
	ret := ValidateGetVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("must be a valid UUID", errs["id"].Error())
}
