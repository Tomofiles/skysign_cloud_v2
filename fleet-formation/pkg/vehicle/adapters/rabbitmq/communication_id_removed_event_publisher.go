package rabbitmq

import (
	"fleet-formation/pkg/vehicle/domain/vehicle"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const communicationIDRemovedEventExchangeName = "vehicle.communication_id_removed_event"

// PublishCommunicationIDRemovedEvent .
func PublishCommunicationIDRemovedEvent(
	ch crm.Channel,
	event vehicle.CommunicationIDRemovedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		communicationIDRemovedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.CommunicationIdRemovedEvent{
		CommunicationId: event.GetCommunicationID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		communicationIDRemovedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", communicationIDRemovedEventExchangeName, eventPb.String())
	return nil
}
