package service

import (
	"context"
	"flightreport/pkg/flightreport/domain/event"
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"flightreport/pkg/flightreport/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightreportID = frep.ID("flightreport-id")
const DefaultFlightoperationID = frep.FlightoperationID("flightoperation-id")

type flightreportRepositoryMock struct {
	mock.Mock
}

func (r *flightreportRepositoryMock) GetAll(
	tx txmanager.Tx,
) ([]*frep.Flightreport, error) {
	ret := r.Called()
	var f []*frep.Flightreport
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*frep.Flightreport)
	}
	return f, ret.Error(1)
}

func (r *flightreportRepositoryMock) GetByID(
	tx txmanager.Tx,
	id frep.ID,
) (*frep.Flightreport, error) {
	ret := r.Called(id)
	var f *frep.Flightreport
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*frep.Flightreport)
	}
	return f, ret.Error(1)
}

func (r *flightreportRepositoryMock) Save(
	tx txmanager.Tx,
	flightreport *frep.Flightreport,
) error {
	ret := r.Called(flightreport)
	return ret.Error(0)
}

type generatorMockFlightreport struct {
	frep.Generator
	id                frep.ID
	flightoperationID frep.FlightoperationID
}

func (gen *generatorMockFlightreport) NewID() frep.ID {
	return gen.id
}
func (gen *generatorMockFlightreport) NewFlightoperationID() frep.FlightoperationID {
	return gen.flightoperationID
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

type flightreportComponentMock struct {
	ID                string
	FlightoperationID string
}

func (f *flightreportComponentMock) GetID() string {
	return f.ID
}

func (f *flightreportComponentMock) GetFlightoperationID() string {
	return f.FlightoperationID
}

type flightreportIDRequestMock struct {
	ID string
}

func (f *flightreportIDRequestMock) GetID() string {
	return f.ID
}

type flightoperationIDRequestMock struct {
	FlightoperationID string
}

func (f *flightoperationIDRequestMock) GetFlightoperationID() string {
	return f.FlightoperationID
}
