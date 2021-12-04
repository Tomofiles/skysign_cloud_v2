package proto

// import (
// 	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
// 	validation "github.com/go-ozzo/ozzo-validation"
// 	"github.com/go-ozzo/ozzo-validation/is"
// )

// // ValidateUpdateMissionRequest .
// func ValidateUpdateMissionRequest(request *skysign_proto.Mission) error {
// 	return validation.ValidateStruct(request,
// 		validation.Field(&request.Id, validation.Required, validation.Length(36, 36), is.UUID),
// 		validation.Field(&request.Name, validation.Required, validation.Length(0, 200)),
// 		validation.Field(&request.Navigation, validation.Required),
// 	)
// }
