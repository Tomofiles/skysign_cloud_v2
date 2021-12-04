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

// TestValidateMissionCopiedEvent_Success .
func TestValidateMissionCopiedEvent_Success(t *testing.T) {
	a := assert.New(t)

	fleetID, _ := uuid.NewRandom()
	originalMissionID, _ := uuid.NewRandom()
	newMissionID, _ := uuid.NewRandom()

	event := &skysign_proto.MissionCopiedEvent{
		FleetId:           fleetID.String(),
		OriginalMissionId: originalMissionID.String(),
		NewMissionId:      newMissionID.String(),
	}
	ret := ValidateMissionCopiedEvent(event)

	a.Nil(ret)
}

// TestValidateMissionCopiedEvent_Failure_Blank .
func TestValidateMissionCopiedEvent_Failure_Blank(t *testing.T) {
	a := assert.New(t)

	fleetID := strings.Repeat("X", 0)
	originalMissionID := strings.Repeat("X", 0)
	newMissionID := strings.Repeat("X", 0)

	event := &skysign_proto.MissionCopiedEvent{
		FleetId:           fleetID,
		OriginalMissionId: originalMissionID,
		NewMissionId:      newMissionID,
	}
	ret := ValidateMissionCopiedEvent(event)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("cannot be blank", errs["fleet_id"].Error())
	a.Equal("cannot be blank", errs["original_mission_id"].Error())
	a.Equal("cannot be blank", errs["new_mission_id"].Error())
}

// TestValidateMissionCopiedEvent_Failure_Length .
func TestValidateMissionCopiedEvent_Failure_Length(t *testing.T) {
	a := assert.New(t)

	fleetID := strings.Repeat("X", 37)
	originalMissionID := strings.Repeat("X", 37)
	newMissionID := strings.Repeat("X", 37)

	event := &skysign_proto.MissionCopiedEvent{
		FleetId:           fleetID,
		OriginalMissionId: originalMissionID,
		NewMissionId:      newMissionID,
	}
	ret := ValidateMissionCopiedEvent(event)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("the length must be exactly 36", errs["fleet_id"].Error())
	a.Equal("the length must be exactly 36", errs["original_mission_id"].Error())
	a.Equal("the length must be exactly 36", errs["new_mission_id"].Error())
}

// TestValidateMissionCopiedEvent_Failure_IllegalFormat .
func TestValidateMissionCopiedEvent_Failure_IllegalFormat(t *testing.T) {
	a := assert.New(t)

	fleetID := strings.Repeat("X", 36)
	originalMissionID := strings.Repeat("X", 36)
	newMissionID := strings.Repeat("X", 36)

	event := &skysign_proto.MissionCopiedEvent{
		FleetId:           fleetID,
		OriginalMissionId: originalMissionID,
		NewMissionId:      newMissionID,
	}
	ret := ValidateMissionCopiedEvent(event)

	var errs validation.Errors
	errors.As(ret, &errs)

	a.Len(errs, 3)
	a.Equal("must be a valid UUID", errs["fleet_id"].Error())
	a.Equal("must be a valid UUID", errs["original_mission_id"].Error())
	a.Equal("must be a valid UUID", errs["new_mission_id"].Error())
}
