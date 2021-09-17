package rabbitmq

import (
	"remote-communication/pkg/communication/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

// TestSubscribeEventHandlerCommunicationIdGaveEvent .
func TestSubscribeEventHandlerCommunicationIdGaveEvent(t *testing.T) {
	a := assert.New(t)

	service := manageCommunicationServiceMock{}
	service.On("CreateCommunication", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageCommunication: &service,
		},
	}

	psm := &publishHandlerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.CommunicationIdGaveEvent{
		CommunicationId: string(DefaultCommunicationID),
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "vehicle.communication_id_gave_event"
		QueueName    = "communication.communication_id_gave_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.ID, string(DefaultCommunicationID))
}

// TestSubscribeEventHandlerCommunicationIdRemovedEvent .
func TestSubscribeEventHandlerCommunicationIdRemovedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageCommunicationServiceMock{}
	service.On("DeleteCommunication", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageCommunication: &service,
		},
	}

	psm := &publishHandlerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.CommunicationIdRemovedEvent{
		CommunicationId: string(DefaultCommunicationID),
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "vehicle.communication_id_removed_event"
		QueueName    = "communication.communication_id_removed_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.ID, string(DefaultCommunicationID))
}
