package rabbitmq

import (
	fope "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/domain/flightoperation"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightoperationCompletedEventExchangeName = "flightoperation.flightoperation_completed_event"

// PublishFlightoperationCompletedEvent .
func PublishFlightoperationCompletedEvent(
	ch crm.Channel,
	event fope.FlightoperationCompletedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightoperationCompletedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightoperationCompletedEvent{
		FlightoperationId: event.GetID(),
		Flightoperation: &skysign_proto.Flightoperation{
			Id:          event.GetID(),
			Name:        event.GetName(),
			Description: event.GetDescription(),
			FleetId:     event.GetFleetID(),
		},
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightoperationCompletedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightoperationCompletedEventExchangeName, eventPb.String())
	return nil
}
