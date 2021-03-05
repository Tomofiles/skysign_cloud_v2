package service

import (
	"context"
	"flightoperation/pkg/flightoperation/domain/event"
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
	"flightoperation/pkg/flightoperation/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightoperationID = fope.ID("flightoperation-id")
const DefaultFlightplanID = fope.FlightplanID("flightplan-id")

type flightoperationRepositoryMock struct {
	mock.Mock
}

func (r *flightoperationRepositoryMock) GetAll(
	tx txmanager.Tx,
) ([]*fope.Flightoperation, error) {
	ret := r.Called()
	var f []*fope.Flightoperation
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*fope.Flightoperation)
	}
	return f, ret.Error(1)
}

func (r *flightoperationRepositoryMock) GetByID(
	tx txmanager.Tx,
	id fope.ID,
) (*fope.Flightoperation, error) {
	ret := r.Called(id)
	var f *fope.Flightoperation
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*fope.Flightoperation)
	}
	return f, ret.Error(1)
}

func (r *flightoperationRepositoryMock) Save(
	tx txmanager.Tx,
	flightoperation *fope.Flightoperation,
) error {
	ret := r.Called(flightoperation)
	return ret.Error(0)
}

type generatorMockFlightoperation struct {
	fope.Generator
	id           fope.ID
	flightplanID fope.FlightplanID
}

func (gen *generatorMockFlightoperation) NewID() fope.ID {
	return gen.id
}
func (gen *generatorMockFlightoperation) NewFlightplanID() fope.FlightplanID {
	return gen.flightplanID
}

type publisherMock struct {
	events  []interface{}
	isFlush bool
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	rm.isFlush = true
	return nil
}

type pubSubManagerMock struct {
	mock.Mock
}

func (psm *pubSubManagerMock) GetPublisher() (event.Publisher, event.ChannelClose, error) {
	ret := psm.Called()
	var pub event.Publisher
	var close func() error
	if ret.Get(0) == nil {
		pub = nil
	} else {
		pub = ret.Get(0).(event.Publisher)
		close = ret.Get(1).(func() error)
	}
	return pub, close, ret.Error(2)
}

func (psm *pubSubManagerMock) SetConsumer(ctx context.Context, exchangeName string, handler event.Handler) error {
	return nil
}

type txManagerMock struct {
	isOpe, isEH error
}

func (txm *txManagerMock) Do(operation func(txmanager.Tx) error) error {
	txm.isOpe = operation(nil)
	return nil
}

func (txm *txManagerMock) DoAndEndHook(operation func(txmanager.Tx) error, endHook func() error) error {
	txm.isOpe = operation(nil)
	txm.isEH = endHook()
	return nil
}

type flightoperationComponentMock struct {
	ID           string
	FlightplanID string
}

func (f *flightoperationComponentMock) GetID() string {
	return f.ID
}

func (f *flightoperationComponentMock) GetFlightplanID() string {
	return f.FlightplanID
}

type flightoperationIDRequestMock struct {
	ID string
}

func (f *flightoperationIDRequestMock) GetID() string {
	return f.ID
}

type flightplanIDRequestMock struct {
	FlightplanID string
}

func (f *flightplanIDRequestMock) GetFlightplanID() string {
	return f.FlightplanID
}
