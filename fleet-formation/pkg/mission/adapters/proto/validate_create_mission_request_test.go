package proto

import (
	"errors"
	"strings"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/stretchr/testify/assert"
)

// TestValidateCreateMissionRequest_Success .
func TestValidateCreateMissionRequest_Success(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Name: name,
		Navigation: &skysign_proto.Navigation{
			TakeoffPointGroundAltitude: 10,
			Waypoints: []*skysign_proto.Waypoint{
				{
					Latitude:         -90.0,
					Longitude:        -180.0,
					RelativeAltitude: 30,
					Speed:            0.1,
				},
				{
					Latitude:         90.0,
					Longitude:        180.0,
					RelativeAltitude: 30,
					Speed:            1,
				},
			},
		},
	}
	ret := ValidateCreateMissionRequest(req)

	a.Nil(ret)
}

// TestValidateCreateMissionRequest_Failure_Blank_FirstLayer .
func TestValidateCreateMissionRequest_Failure_Blank_FirstLayer(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 0)

	req := &skysign_proto.Mission{
		Name: name,
	}
	ret := ValidateCreateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 2)
	a.Equal("cannot be blank", errs["name"].Error())
	a.Equal("cannot be blank", errs["navigation"].Error())
}

// TestValidateCreateMissionRequest_Failure_Blank_SecondLayer .
func TestValidateCreateMissionRequest_Failure_Blank_SecondLayer(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Name:       name,
		Navigation: &skysign_proto.Navigation{},
	}
	ret := ValidateCreateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("cannot be blank", errs["waypoints"].Error())
}

// TestValidateCreateMissionRequest_Failure_Length .
func TestValidateCreateMissionRequest_Failure_Length(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 201)

	req := &skysign_proto.Mission{
		Name:       name,
		Navigation: &skysign_proto.Navigation{},
	}
	ret := ValidateCreateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("the length must be no more than 200", errs["name"].Error())
}

// TestValidateCreateMissionRequest_Failure_Min .
func TestValidateCreateMissionRequest_Failure_Min(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Name: name,
		Navigation: &skysign_proto.Navigation{
			TakeoffPointGroundAltitude: 10,
			Waypoints: []*skysign_proto.Waypoint{
				{
					Latitude:         -90.1,
					Longitude:        -180.1,
					RelativeAltitude: 10,
					Speed:            0.01,
				},
			},
		},
	}
	ret := ValidateCreateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("must be no less than -90", errs["latitude"].Error())
	a.Equal("must be no less than -180", errs["longitude"].Error())
	a.Equal("must be no less than 0.1", errs["speed"].Error())
}

// TestValidateCreateMissionRequest_Failure_Max .
func TestValidateCreateMissionRequest_Failure_Max(t *testing.T) {
	a := assert.New(t)

	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Name: name,
		Navigation: &skysign_proto.Navigation{
			TakeoffPointGroundAltitude: 10,
			Waypoints: []*skysign_proto.Waypoint{
				{
					Latitude:         90.1,
					Longitude:        180.1,
					RelativeAltitude: 10,
					Speed:            10,
				},
			},
		},
	}
	ret := ValidateCreateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 2)
	a.Equal("must be no greater than 90", errs["latitude"].Error())
	a.Equal("must be no greater than 180", errs["longitude"].Error())
}
