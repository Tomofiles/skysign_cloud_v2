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

// TestValidateDeleteMissionRequest_Success .
func TestValidateDeleteMissionRequest_Success(t *testing.T) {
	a := assert.New(t)

	id, _ := uuid.NewRandom()

	req := &skysign_proto.DeleteMissionRequest{
		Id: id.String(),
	}
	ret := ValidateDeleteMissionRequest(req)

	a.Nil(ret)
}

// TestValidateDeleteMissionRequest_Failure_Blank .
func TestValidateDeleteMissionRequest_Failure_Blank(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 0)

	req := &skysign_proto.DeleteMissionRequest{
		Id: id,
	}
	ret := ValidateDeleteMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("cannot be blank", errs["id"].Error())
}

// TestValidateDeleteMissionRequest_Failure_Length .
func TestValidateDeleteMissionRequest_Failure_Length(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 37)

	req := &skysign_proto.DeleteMissionRequest{
		Id: id,
	}
	ret := ValidateDeleteMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("the length must be exactly 36", errs["id"].Error())
}

// TestValidateDeleteMissionRequest_Failure_IllegalFormat .
func TestValidateDeleteMissionRequest_Failure_IllegalFormat(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 36)

	req := &skysign_proto.DeleteMissionRequest{
		Id: id,
	}
	ret := ValidateDeleteMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("must be a valid UUID", errs["id"].Error())
}
