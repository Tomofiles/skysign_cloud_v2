package proto

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// ValidateGetVehicleRequest .
func ValidateGetVehicleRequest(request *skysign_proto.GetVehicleRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required, validation.Length(36, 36), is.UUID),
	)
}

// func (f *updateCommand) Validation() error {
// 	return validation.ValidateStruct(f.request,
// 		validation.Field(&f.request.Id, validation.Required, validation.Length(36, 36), is.UUID),
// 		validation.Field(&f.request.Name, validation.Required, validation.Length(1, 200)),
// 		validation.Field(&f.request.CommunicationId, validation.Required, validation.Length(1, 36)),
// 	)
// }
