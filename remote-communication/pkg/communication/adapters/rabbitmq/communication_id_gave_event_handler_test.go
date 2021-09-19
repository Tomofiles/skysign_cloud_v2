package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleCommunicationIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	service := manageCommunicationServiceMock{}

	service.On("CreateCommunication", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageCommunication: &service,
		},
	}

	handler := NewCommunicationIDGaveEventHandler(app)

	requestPb := &skysign_proto.CommunicationIdGaveEvent{
		CommunicationId: string(DefaultCommunicationID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleCommunicationIDGaveEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.ID, string(DefaultCommunicationID))
}
