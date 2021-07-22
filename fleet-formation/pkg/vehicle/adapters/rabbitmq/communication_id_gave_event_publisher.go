package rabbitmq

import (
	crm "fleet-formation/pkg/common/adapters/rabbitmq"
	"fleet-formation/pkg/skysign_proto"
	"fleet-formation/pkg/vehicle/domain/vehicle"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const communicationIDGaveEventExchangeName = "vehicle.communication_id_gave_event"

// PublishCommunicationIDGaveEvent .
func PublishCommunicationIDGaveEvent(
	ch crm.Channel,
	event vehicle.CommunicationIDGaveEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		communicationIDGaveEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.CommunicationIdGaveEvent{
		CommunicationId: event.GetCommunicationID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		communicationIDGaveEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", communicationIDGaveEventExchangeName, eventPb.String())
	return nil
}
