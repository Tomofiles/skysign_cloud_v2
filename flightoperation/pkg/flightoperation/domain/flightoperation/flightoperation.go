package flightoperation

import (
	"errors"
	"flightoperation/pkg/flightoperation/domain/event"
)

// ID .
type ID string

// FlightplanID .
type FlightplanID string

// Version .
type Version string

const (
	// Operating .
	Operating = false
	// Completed .
	Completed = true
)

var (
	// ErrCannotChange .
	ErrCannotChange = errors.New("cannnot change completed flightoperation")
)

// Flightoperation .
type Flightoperation struct {
	id           ID
	flightplanID FlightplanID
	isCompleted  bool
	version      Version
	newVersion   Version
	gen          Generator
	pub          event.Publisher
}

// SetPublisher .
func (f *Flightoperation) SetPublisher(pub event.Publisher) {
	f.pub = pub
}

// GetID .
func (f *Flightoperation) GetID() ID {
	return f.id
}

// GetFlightplanID .
func (f *Flightoperation) GetFlightplanID() FlightplanID {
	return f.flightplanID
}

// GetVersion .
func (f *Flightoperation) GetVersion() Version {
	return f.version
}

// GetNewVersion .
func (f *Flightoperation) GetNewVersion() Version {
	return f.newVersion
}

// Complete .
func (f *Flightoperation) Complete() error {
	if f.isCompleted {
		return ErrCannotChange
	}
	if f.pub != nil {
		f.pub.Publish(CompletedEvent{
			ID: f.id,
		})
	}
	f.isCompleted = Completed
	f.newVersion = f.gen.NewVersion()

	return nil
}

// Generator .
type Generator interface {
	NewID() ID
	NewFlightplanID() FlightplanID
	NewVersion() Version
}
