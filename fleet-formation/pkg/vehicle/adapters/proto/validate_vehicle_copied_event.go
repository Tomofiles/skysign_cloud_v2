package proto

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// ValidateVehicleCopiedEvent .
func ValidateVehicleCopiedEvent(event *skysign_proto.VehicleCopiedEvent) error {
	return validation.ValidateStruct(event,
		validation.Field(&event.FleetId, validation.Required, validation.Length(36, 36), is.UUID),
		validation.Field(&event.OriginalVehicleId, validation.Required, validation.Length(36, 36), is.UUID),
		validation.Field(&event.NewVehicleId, validation.Required, validation.Length(36, 36), is.UUID),
	)
}
