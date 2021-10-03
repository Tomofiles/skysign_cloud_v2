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

// TestValidateDeleteVehicleRequest_Success .
func TestValidateDeleteVehicleRequest_Success(t *testing.T) {
	a := assert.New(t)

	id, _ := uuid.NewRandom()

	req := &skysign_proto.DeleteVehicleRequest{
		Id: id.String(),
	}
	ret := ValidateDeleteVehicleRequest(req)

	a.Nil(ret)
}

// TestValidateDeleteVehicleRequest_Failure_Blank .
func TestValidateDeleteVehicleRequest_Failure_Blank(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 0)

	req := &skysign_proto.DeleteVehicleRequest{
		Id: id,
	}
	ret := ValidateDeleteVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("cannot be blank", errs["id"].Error())
}

// TestValidateDeleteVehicleRequest_Failure_Length .
func TestValidateDeleteVehicleRequest_Failure_Length(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 37)

	req := &skysign_proto.DeleteVehicleRequest{
		Id: id,
	}
	ret := ValidateDeleteVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("the length must be exactly 36", errs["id"].Error())
}

// TestValidateDeleteVehicleRequest_Failure_IllegalFormat .
func TestValidateDeleteVehicleRequest_Failure_IllegalFormat(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 36)

	req := &skysign_proto.DeleteVehicleRequest{
		Id: id,
	}
	ret := ValidateDeleteVehicleRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("must be a valid UUID", errs["id"].Error())
}
