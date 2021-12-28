package proto

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// ValidateUpdateMissionRequest .
func ValidateUpdateMissionRequest(request *skysign_proto.Mission) error {
	if err := validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required, validation.Length(36, 36), is.UUID),
		validation.Field(&request.Name, validation.Required, validation.Length(0, 200)),
		validation.Field(&request.Navigation, validation.Required),
	); err != nil {
		return err
	}
	navigation := request.Navigation
	if err := validation.ValidateStruct(navigation,
		validation.Field(&navigation.TakeoffPointGroundAltitude),
		validation.Field(&navigation.Waypoints, validation.Required),
	); err != nil {
		return err
	}
	for _, waypoint := range navigation.Waypoints {
		if err := validation.ValidateStruct(waypoint,
			validation.Field(&waypoint.Latitude, validation.Min(-90.0), validation.Max(90.0)),
			validation.Field(&waypoint.Longitude, validation.Min(-180.0), validation.Max(180.0)),
			validation.Field(&waypoint.RelativeAltitude),
			validation.Field(&waypoint.Speed, validation.Min(0.1)),
		); err != nil {
			return err
		}
	}
	return nil
}
