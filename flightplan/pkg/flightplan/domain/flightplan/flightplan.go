package flightplan

import "errors"

// ID .
type ID string

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
	ErrCannotChange = errors.New("cannnot carbon copied flightplan")
)

// Flightplan .
type Flightplan struct {
	id           ID
	name         string
	description  string
	isCarbonCopy bool
	version      Version
	newVersion   Version
	gen          Generator
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
	if f.isCarbonCopy {
		return ErrCannotChange
	}
	f.name = name
	f.newVersion = f.gen.NewVersion()
	return nil
}

// ChangeDescription .
func (f *Flightplan) ChangeDescription(description string) error {
	if f.isCarbonCopy {
		return ErrCannotChange
	}
	f.description = description
	f.newVersion = f.gen.NewVersion()
	return nil
}

// Generator .
type Generator interface {
	NewID() ID
	NewVersion() Version
}
