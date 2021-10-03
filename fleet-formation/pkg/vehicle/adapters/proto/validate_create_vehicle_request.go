package proto

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
)

// ValidateCreateVehicleRequest .
func ValidateCreateVehicleRequest(request *skysign_proto.Vehicle) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Name, validation.Required, validation.Length(0, 200)),
		validation.Field(&request.CommunicationId, validation.Required, validation.Length(0, 36)),
	)
}
