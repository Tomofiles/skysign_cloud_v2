package flightoperation

import (
	"errors"
	"flightoperation/pkg/flightoperation/domain/event"
)

// ID .
type ID string

// FleetID .
type FleetID string

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
	id          ID
	name        string
	description string
	fleetID     FleetID
	isCompleted bool
	version     Version
	newVersion  Version
	gen         Generator
	pub         event.Publisher
}

// SetPublisher .
func (f *Flightoperation) SetPublisher(pub event.Publisher) {
	f.pub = pub
}

// GetID .
func (f *Flightoperation) GetID() ID {
	return f.id
}

// GetName .
func (f *Flightoperation) GetName() string {
	return f.name
}

// GetDescription .
func (f *Flightoperation) GetDescription() string {
	return f.description
}

// GetFleetID .
func (f *Flightoperation) GetFleetID() FleetID {
	return f.fleetID
}

// GetVersion .
func (f *Flightoperation) GetVersion() Version {
	return f.version
}

// GetNewVersion .
func (f *Flightoperation) GetNewVersion() Version {
	return f.newVersion
}

// NameFlightoperation .
func (f *Flightoperation) NameFlightoperation(name string) error {
	f.name = name
	f.newVersion = f.gen.NewVersion()
	return nil
}

// ChangeDescription .
func (f *Flightoperation) ChangeDescription(description string) error {
	f.description = description
	f.newVersion = f.gen.NewVersion()
	return nil
}

// Complete .
func (f *Flightoperation) Complete() error {
	if f.isCompleted {
		return ErrCannotChange
	}
	if f.pub != nil {
		f.pub.Publish(FlightoperationCompletedEvent{
			ID:          f.id,
			Name:        f.name,
			Description: f.description,
			FleetID:     f.fleetID,
		})
	}
	f.isCompleted = Completed
	f.newVersion = f.gen.NewVersion()

	return nil
}

// Generator .
type Generator interface {
	NewID() ID
	NewFleetID() FleetID
	NewVersion() Version
}
