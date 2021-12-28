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

// TestValidateUpdateMissionRequest_Success .
func TestValidateUpdateMissionRequest_Success(t *testing.T) {
	a := assert.New(t)

	uuid, _ := uuid.NewRandom()

	id := uuid.String()
	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Id:   id,
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
	ret := ValidateUpdateMissionRequest(req)

	a.Nil(ret)
}

// TestValidateUpdateMissionRequest_Failure_Blank_FirstLayer .
func TestValidateUpdateMissionRequest_Failure_Blank_FirstLayer(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 0)
	name := strings.Repeat("X", 0)

	req := &skysign_proto.Mission{
		Id:   id,
		Name: name,
	}
	ret := ValidateUpdateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("cannot be blank", errs["id"].Error())
	a.Equal("cannot be blank", errs["name"].Error())
	a.Equal("cannot be blank", errs["navigation"].Error())
}

// TestValidateUpdateMissionRequest_Failure_Blank_SecondLayer .
func TestValidateUpdateMissionRequest_Failure_Blank_SecondLayer(t *testing.T) {
	a := assert.New(t)

	uuid, _ := uuid.NewRandom()

	id := uuid.String()
	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Id:         id,
		Name:       name,
		Navigation: &skysign_proto.Navigation{},
	}
	ret := ValidateUpdateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 1)
	a.Equal("cannot be blank", errs["waypoints"].Error())
}

// TestValidateUpdateMissionRequest_Failure_Length .
func TestValidateUpdateMissionRequest_Failure_Length(t *testing.T) {
	a := assert.New(t)

	id := strings.Repeat("X", 37)
	name := strings.Repeat("X", 201)

	req := &skysign_proto.Mission{
		Id:         id,
		Name:       name,
		Navigation: &skysign_proto.Navigation{},
	}
	ret := ValidateUpdateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 2)
	a.Equal("the length must be exactly 36", errs["id"].Error())
	a.Equal("the length must be no more than 200", errs["name"].Error())
}

// TestValidateUpdateMissionRequest_Failure_Min .
func TestValidateUpdateMissionRequest_Failure_Min(t *testing.T) {
	a := assert.New(t)

	uuid, _ := uuid.NewRandom()

	id := uuid.String()
	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Id:   id,
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
	ret := ValidateUpdateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("must be no less than -90", errs["latitude"].Error())
	a.Equal("must be no less than -180", errs["longitude"].Error())
	a.Equal("must be no less than 0.1", errs["speed"].Error())
}

// TestValidateUpdateMissionRequest_Failure_Max .
func TestValidateUpdateMissionRequest_Failure_Max(t *testing.T) {
	a := assert.New(t)

	uuid, _ := uuid.NewRandom()

	id := uuid.String()
	name := strings.Repeat("X", 200)

	req := &skysign_proto.Mission{
		Id:   id,
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
	ret := ValidateUpdateMissionRequest(req)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 2)
	a.Equal("must be no greater than 90", errs["latitude"].Error())
	a.Equal("must be no greater than 180", errs["longitude"].Error())
}
