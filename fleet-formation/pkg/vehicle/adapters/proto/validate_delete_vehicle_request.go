package proto

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// ValidateDeleteVehicleRequest .
func ValidateDeleteVehicleRequest(request *skysign_proto.DeleteVehicleRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required, validation.Length(36, 36), is.UUID),
	)
}
