package proto

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// ValidateUpdateVehicleRequest .
func ValidateUpdateVehicleRequest(request *skysign_proto.Vehicle) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required, validation.Length(36, 36), is.UUID),
		validation.Field(&request.Name, validation.Required, validation.Length(0, 200)),
		validation.Field(&request.CommunicationId, validation.Required, validation.Length(0, 36)),
	)
}
