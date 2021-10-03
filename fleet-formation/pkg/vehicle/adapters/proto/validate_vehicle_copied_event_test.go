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

// TestValidateVehicleCopiedEvent_Success .
func TestValidateVehicleCopiedEvent_Success(t *testing.T) {
	a := assert.New(t)

	fleetID, _ := uuid.NewRandom()
	originalVehicleID, _ := uuid.NewRandom()
	newVehicleID, _ := uuid.NewRandom()

	event := &skysign_proto.VehicleCopiedEvent{
		FleetId:           fleetID.String(),
		OriginalVehicleId: originalVehicleID.String(),
		NewVehicleId:      newVehicleID.String(),
	}
	ret := ValidateVehicleCopiedEvent(event)

	a.Nil(ret)
}

// TestValidateVehicleCopiedEvent_Failure_Blank .
func TestValidateVehicleCopiedEvent_Failure_Blank(t *testing.T) {
	a := assert.New(t)

	fleetID := strings.Repeat("X", 0)
	originalVehicleID := strings.Repeat("X", 0)
	newVehicleID := strings.Repeat("X", 0)

	event := &skysign_proto.VehicleCopiedEvent{
		FleetId:           fleetID,
		OriginalVehicleId: originalVehicleID,
		NewVehicleId:      newVehicleID,
	}
	ret := ValidateVehicleCopiedEvent(event)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("cannot be blank", errs["fleet_id"].Error())
	a.Equal("cannot be blank", errs["original_vehicle_id"].Error())
	a.Equal("cannot be blank", errs["new_vehicle_id"].Error())
}

// TestValidateVehicleCopiedEvent_Failure_Length .
func TestValidateVehicleCopiedEvent_Failure_Length(t *testing.T) {
	a := assert.New(t)

	fleetID := strings.Repeat("X", 37)
	originalVehicleID := strings.Repeat("X", 37)
	newVehicleID := strings.Repeat("X", 37)

	event := &skysign_proto.VehicleCopiedEvent{
		FleetId:           fleetID,
		OriginalVehicleId: originalVehicleID,
		NewVehicleId:      newVehicleID,
	}
	ret := ValidateVehicleCopiedEvent(event)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("the length must be exactly 36", errs["fleet_id"].Error())
	a.Equal("the length must be exactly 36", errs["original_vehicle_id"].Error())
	a.Equal("the length must be exactly 36", errs["new_vehicle_id"].Error())
}

// TestValidateVehicleCopiedEvent_Failure_IllegalFormat .
func TestValidateVehicleCopiedEvent_Failure_IllegalFormat(t *testing.T) {
	a := assert.New(t)

	fleetID := strings.Repeat("X", 36)
	originalVehicleID := strings.Repeat("X", 36)
	newVehicleID := strings.Repeat("X", 36)

	event := &skysign_proto.VehicleCopiedEvent{
		FleetId:           fleetID,
		OriginalVehicleId: originalVehicleID,
		NewVehicleId:      newVehicleID,
	}
	ret := ValidateVehicleCopiedEvent(event)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("must be a valid UUID", errs["fleet_id"].Error())
	a.Equal("must be a valid UUID", errs["original_vehicle_id"].Error())
	a.Equal("must be a valid UUID", errs["new_vehicle_id"].Error())
}
