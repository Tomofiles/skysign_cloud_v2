package ports

import (
	"context"
	"remote-communication/pkg/communication/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// CommunicationIDGaveEventExchangeName .
	CommunicationIDGaveEventExchangeName = "vehicle.communication_id_gave_event"
	// CommunicationIDGaveEventQueueName .
	CommunicationIDGaveEventQueueName = "communication.communication_id_gave_event"
	// CommunicationIDRemovedEventExchangeName .
	CommunicationIDRemovedEventExchangeName = "vehicle.communication_id_removed_event"
	// CommunicationIDRemovedEventQueueName .
	CommunicationIDRemovedEventQueueName = "communication.communication_id_removed_event"
)

// EventHandler .
type EventHandler interface {
	HandleCommunicationIDGaveEvent(ctx context.Context, event []byte) error
	HandleCommunicationIDRemovedEvent(ctx context.Context, event []byte) error
}

// eventHandler .
type eventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) *eventHandler {
	return &eventHandler{app: application}
}

// HandleCommunicationIDGaveEvent .
func (h *eventHandler) HandleCommunicationIDGaveEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.CommunicationIdGaveEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", CommunicationIDGaveEventQueueName, eventPb.String())

	command := createCommunicationCommandHolder{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageCommunication.CreateCommunication(&command); ret != nil {
		return ret
	}
	return nil
}

// HandleCommunicationIDRemovedEvent .
func (h *eventHandler) HandleCommunicationIDRemovedEvent(
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

type createCommunicationCommandHolder struct {
	event *skysign_proto.CommunicationIdGaveEvent
}

func (h *createCommunicationCommandHolder) GetID() string {
	return h.event.CommunicationId
}

type deleteCommunicationCommandHolder struct {
	event *skysign_proto.CommunicationIdRemovedEvent
}

func (h *deleteCommunicationCommandHolder) GetID() string {
	return h.event.CommunicationId
}
