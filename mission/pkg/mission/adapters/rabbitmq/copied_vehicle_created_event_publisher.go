package rabbitmq

// import (
// 	"vehicle/pkg/skysign_proto"
// 	"vehicle/pkg/vehicle/domain/vehicle"

// 	"github.com/golang/glog"
// 	"google.golang.org/protobuf/proto"
// )

// const copiedVehicleCreatedEventExchangeName = "vehicle.copied_vehicle_created_event"

// // PublishCopiedVehicleCreatedEvent .
// func PublishCopiedVehicleCreatedEvent(
// 	ch Channel,
// 	event vehicle.CopiedVehicleCreatedEvent,
// ) error {
// 	if err := ch.FanoutExchangeDeclare(
// 		copiedVehicleCreatedEventExchangeName,
// 	); err != nil {
// 		return err
// 	}

// 	eventPb := skysign_proto.CopiedVehicleCreatedEvent{
// 		VehicleId:       event.GetVehicleID(),
// 		CommunicationId: event.GetCommunicationID(),
// 		FlightplanId:    event.GetFlightplanID(),
// 	}
// 	eventBin, err := proto.Marshal(&eventPb)
// 	if err != nil {
// 		return err
// 	}

// 	if err := ch.Publish(
// 		copiedVehicleCreatedEventExchangeName,
// 		eventBin,
// 	); err != nil {
// 		return err
// 	}

// 	glog.Infof("PUBLISH , Event: %s, Message: %s", copiedVehicleCreatedEventExchangeName, eventPb.String())
// 	return nil
// }
