package flightoperation

import "flightoperation/pkg/flightoperation/domain/event"

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
func (f *Flightoperation) Complete() {
	if f.pub != nil && f.isCompleted != Completed {
		f.pub.Publish(&CompletedEvent{
			ID: f.id,
		})
	}
	f.isCompleted = Completed
	f.newVersion = f.gen.NewVersion()
}

// Generator .
type Generator interface {
	NewID() ID
	NewFlightplanID() FlightplanID
	NewVersion() Version
}
