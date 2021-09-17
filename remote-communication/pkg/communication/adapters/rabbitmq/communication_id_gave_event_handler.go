package rabbitmq

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
)

// CommunicationIDGaveEventHandler .
type CommunicationIDGaveEventHandler interface {
	HandleCommunicationIDGaveEvent(ctx context.Context, event []byte) error
}

// communicationIDGaveEventHandler .
type communicationIDGaveEventHandler struct {
	app app.Application
}

// NewCommunicationIDGaveEventHandler .
func NewCommunicationIDGaveEventHandler(application app.Application) *communicationIDGaveEventHandler {
	return &communicationIDGaveEventHandler{app: application}
}

// HandleCommunicationIDGaveEvent .
func (h *communicationIDGaveEventHandler) HandleCommunicationIDGaveEvent(
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

type createCommunicationCommandHolder struct {
	event *skysign_proto.CommunicationIdGaveEvent
}

func (h *createCommunicationCommandHolder) GetID() string {
	return h.event.CommunicationId
}
