package proto

// import (
// 	"errors"
// 	"strings"
// 	"testing"

// 	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
// 	validation "github.com/go-ozzo/ozzo-validation"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// // TestValidateUpdateVehicleRequest_Success .
// func TestValidateUpdateVehicleRequest_Success(t *testing.T) {
// 	a := assert.New(t)

// 	id, _ := uuid.NewRandom()
// 	name := strings.Repeat("X", 200)
// 	communicationID := strings.Repeat("X", 36)

// 	req := &skysign_proto.Vehicle{
// 		Id:              id.String(),
// 		Name:            name,
// 		CommunicationId: communicationID,
// 	}
// 	ret := ValidateUpdateVehicleRequest(req)

// 	a.Nil(ret)
// }

// // TestValidateUpdateVehicleRequest_Failure_Blank .
// func TestValidateUpdateVehicleRequest_Failure_Blank(t *testing.T) {
// 	a := assert.New(t)

// 	id := strings.Repeat("X", 0)
// 	name := strings.Repeat("X", 0)
// 	communicationID := strings.Repeat("X", 0)

// 	req := &skysign_proto.Vehicle{
// 		Id:              id,
// 		Name:            name,
// 		CommunicationId: communicationID,
// 	}
// 	ret := ValidateUpdateVehicleRequest(req)

// 	var errs validation.Errors
// 	errors.As(ret, &errs)

// 	a.Len(errs, 3)
// 	a.Equal("cannot be blank", errs["id"].Error())
// 	a.Equal("cannot be blank", errs["name"].Error())
// 	a.Equal("cannot be blank", errs["communication_id"].Error())
// }

// // TestValidateUpdateVehicleRequest_Failure_Length .
// func TestValidateUpdateVehicleRequest_Failure_Length(t *testing.T) {
// 	a := assert.New(t)

// 	id := strings.Repeat("X", 37)
// 	name := strings.Repeat("X", 201)
// 	communicationID := strings.Repeat("X", 37)

// 	req := &skysign_proto.Vehicle{
// 		Id:              id,
// 		Name:            name,
// 		CommunicationId: communicationID,
// 	}
// 	ret := ValidateUpdateVehicleRequest(req)

// 	var errs validation.Errors
// 	errors.As(ret, &errs)

// 	a.Len(errs, 3)
// 	a.Equal("the length must be exactly 36", errs["id"].Error())
// 	a.Equal("the length must be no more than 200", errs["name"].Error())
// 	a.Equal("the length must be no more than 36", errs["communication_id"].Error())
// }

// // TestValidateUpdateVehicleRequest_Failure_IllegalFormat .
// func TestValidateUpdateVehicleRequest_Failure_IllegalFormat(t *testing.T) {
// 	a := assert.New(t)

// 	id := strings.Repeat("X", 36)
// 	name := strings.Repeat("X", 200)
// 	communicationID := strings.Repeat("X", 36)

// 	req := &skysign_proto.Vehicle{
// 		Id:              id,
// 		Name:            name,
// 		CommunicationId: communicationID,
// 	}
// 	ret := ValidateUpdateVehicleRequest(req)

// 	var errs validation.Errors
// 	errors.As(ret, &errs)

// 	a.Len(errs, 1)
// 	a.Equal("must be a valid UUID", errs["id"].Error())
// }
