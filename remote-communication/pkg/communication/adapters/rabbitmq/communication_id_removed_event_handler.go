package rabbitmq

import (
	"context"
	"remote-communication/pkg/communication/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// CommunicationIDRemovedEventExchangeName .
	CommunicationIDRemovedEventExchangeName = "vehicle.communication_id_removed_event"
	// CommunicationIDRemovedEventQueueName .
	CommunicationIDRemovedEventQueueName = "communication.communication_id_removed_event"
)

// CommunicationIDRemovedEventHandler .
type CommunicationIDRemovedEventHandler interface {
	HandleCommunicationIDRemovedEvent(ctx context.Context, event []byte) error
}

// communicationIDRemovedEventHandler .
type communicationIDRemovedEventHandler struct {
	app app.Application
}

// NewCommunicationIDRemovedEventHandler .
func NewCommunicationIDRemovedEventHandler(application app.Application) CommunicationIDRemovedEventHandler {
	return &communicationIDRemovedEventHandler{app: application}
}

// HandleCommunicationIDRemovedEvent .
func (h *communicationIDRemovedEventHandler) HandleCommunicationIDRemovedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.CommunicationIdRemovedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", CommunicationIDRemovedEventQueueName, eventPb.String())

	command := deleteCommunicationCommandHolder{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageCommunication.DeleteCommunication(&command); ret != nil {
		return ret
	}
	return nil
}

type deleteCommunicationCommandHolder struct {
	event *skysign_proto.CommunicationIdRemovedEvent
}

func (h *deleteCommunicationCommandHolder) GetID() string {
	return h.event.CommunicationId
}
