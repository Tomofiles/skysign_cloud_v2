package proto

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// ValidateCreateMissionRequest .
func ValidateCreateMissionRequest(request *skysign_proto.Mission) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Name, validation.Required, validation.Length(0, 200)),
		validation.Field(&request.Navigation, validation.Required,
			ValidateNavigation("upload_id", getUploadIdValue, validation.Required, validation.Length(36, 36), is.UUID),
			ValidateNavigation("takeoff_point_ground_altitude", getTakeoffPointGroundAltitudeValue, validation.Required, validation.Min(1.0)),
		),
	)
}

func getUploadIdValue(value interface{}) interface{} {
	nav, ok := value.(*skysign_proto.Navigation)
	if !ok {
		panic("developer error")
	}
	return nav.UploadId
}

func getTakeoffPointGroundAltitudeValue(value interface{}) interface{} {
	nav, ok := value.(*skysign_proto.Navigation)
	if !ok {
		panic("developer error")
	}
	return nav.TakeoffPointGroundAltitude
}
