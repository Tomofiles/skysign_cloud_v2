package vehicle

import (
	"errors"
	"vehicle/pkg/vehicle/domain/event"
)

// ID .
type ID string

// CommunicationID .
type CommunicationID string

// Version .
type Version string

const (
	// Original .
	Original = false
	// CarbonCopy .
	CarbonCopy = true
)

var (
	// ErrCannotChange .
	ErrCannotChange = errors.New("cannot change carbon copied vehicle")
)

// Vehicle .
type Vehicle struct {
	id              ID
	name            string
	communicationID CommunicationID
	isCarbonCopy    bool
	version         Version
	newVersion      Version
	gen             Generator
	pub             event.Publisher
}

// SetPublisher .
func (v *Vehicle) SetPublisher(pub event.Publisher) {
	v.pub = pub
}

// GetID .
func (v *Vehicle) GetID() ID {
	return v.id
}

// GetName .
func (v *Vehicle) GetName() string {
	return v.name
}

// GetCommunicationID .
func (v *Vehicle) GetCommunicationID() CommunicationID {
	return v.communicationID
}

// GetVersion .
func (v *Vehicle) GetVersion() Version {
	return v.version
}

// GetNewVersion .
func (v *Vehicle) GetNewVersion() Version {
	return v.newVersion
}

// NameVehicle .
func (v *Vehicle) NameVehicle(name string) error {
	if v.isCarbonCopy {
		return ErrCannotChange
	}
	v.name = name
	v.newVersion = v.gen.NewVersion()
	return nil
}

// GiveCommunication .
func (v *Vehicle) GiveCommunication(communicationID CommunicationID) error {
	if v.isCarbonCopy {
		return ErrCannotChange
	}
	if v.pub != nil {
		if v.communicationID == "" {
			v.communicationID = communicationID
			v.newVersion = v.gen.NewVersion()
			v.pub.Publish(CommunicationIdGaveEvent{
				CommunicationID: v.communicationID,
			})
		} else {
			beforeId := v.communicationID
			v.communicationID = communicationID
			v.newVersion = v.gen.NewVersion()
			v.pub.Publish(CommunicationIdRemovedEvent{
				CommunicationID: beforeId,
			})
			v.pub.Publish(CommunicationIdGaveEvent{
				CommunicationID: v.communicationID,
			})
		}
	}
	return nil
}

// RemoveCommunication .
func (v *Vehicle) RemoveCommunication() error {
	if v.isCarbonCopy {
		return ErrCannotChange
	}
	if v.pub != nil {
		v.pub.Publish(CommunicationIdRemovedEvent{
			CommunicationID: v.communicationID,
		})
	}
	v.communicationID = ""
	v.newVersion = v.gen.NewVersion()
	return nil
}

// Generator .
type Generator interface {
	NewID() ID
	NewVersion() Version
}
