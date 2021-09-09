package flightplan

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
)

// ID .
type ID string

// FleetID .
type FleetID string

// Version .
type Version string

const (
	// BlankFleetID .
	BlankFleetID = FleetID("blank-fleet-id-@domain")
)

// Flightplan .
type Flightplan struct {
	id          ID
	name        string
	description string
	fleetID     FleetID
	version     Version
	newVersion  Version
	gen         Generator
	pub         event.Publisher
}

// SetPublisher .
func (f *Flightplan) SetPublisher(pub event.Publisher) {
	f.pub = pub
}

// GetID .
func (f *Flightplan) GetID() ID {
	return f.id
}

// GetName .
func (f *Flightplan) GetName() string {
	return f.name
}

// GetDescription .
func (f *Flightplan) GetDescription() string {
	return f.description
}

// GetFleetID .
func (f *Flightplan) GetFleetID() FleetID {
	return f.fleetID
}

// GetVersion .
func (f *Flightplan) GetVersion() Version {
	return f.version
}

// GetNewVersion .
func (f *Flightplan) GetNewVersion() Version {
	return f.newVersion
}

// NameFlightplan .
func (f *Flightplan) NameFlightplan(name string) error {
	f.name = name
	f.newVersion = f.gen.NewVersion()
	return nil
}

// ChangeDescription .
func (f *Flightplan) ChangeDescription(description string) error {
	f.description = description
	f.newVersion = f.gen.NewVersion()
	return nil
}

// ChangeNumberOfVehicles .
func (f *Flightplan) ChangeNumberOfVehicles(numberOfVehicles int) error {
	if f.fleetID != BlankFleetID {
		if f.pub != nil {
			f.pub.Publish(FleetIDRemovedEvent{
				FleetID: f.fleetID,
			})
		}
	}
	f.fleetID = f.gen.NewFleetID()
	f.newVersion = f.gen.NewVersion()
	if f.pub != nil {
		f.pub.Publish(FleetIDGaveEvent{
			FleetID:          f.fleetID,
			NumberOfVehicles: numberOfVehicles,
		})
	}
	return nil
}

// RemoveFleetID .
func (f *Flightplan) RemoveFleetID() error {
	if f.fleetID != BlankFleetID {
		if f.pub != nil {
			f.pub.Publish(FleetIDRemovedEvent{
				FleetID: f.fleetID,
			})
		}
		f.fleetID = BlankFleetID
		f.newVersion = f.gen.NewVersion()
	}
	return nil
}

// Generator .
type Generator interface {
	NewID() ID
	NewFleetID() FleetID
	NewVersion() Version
}
