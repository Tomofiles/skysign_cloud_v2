package rabbitmq

import (
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/domain/flightplan"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightplanExecutedEventExchangeName = "flightplan.flightplan_executed_event"

// PublishFlightplanExecutedEvent .
func PublishFlightplanExecutedEvent(
	ch crm.Channel,
	event flightplan.FlightplanExecutedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightplanExecutedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightplanExecutedEvent{
		FlightplanId: string(event.GetID()),
		Flightplan: &skysign_proto.Flightplan{
			Id:          string(event.GetID()),
			Name:        event.GetName(),
			Description: event.GetDescription(),
			FleetId:     string(event.GetFleetID()),
		},
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightplanExecutedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightplanExecutedEventExchangeName, eventPb.String())
	return nil
}
