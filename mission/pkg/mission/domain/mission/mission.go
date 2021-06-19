package mission

import (
	"errors"
	"mission/pkg/mission/domain/event"
)

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
	ErrCannotChange = errors.New("cannot change carbon copied mission")
)

// Mission .
type Mission struct {
	id           ID
	name         string
	navigation   *Navigation
	isCarbonCopy bool
	version      Version
	newVersion   Version
	gen          Generator
	pub          event.Publisher
}

// SetPublisher .
func (m *Mission) SetPublisher(pub event.Publisher) {
	m.pub = pub
}

// GetID .
func (m *Mission) GetID() ID {
	return m.id
}

// GetName .
func (m *Mission) GetName() string {
	return m.name
}

// GetNavigation .
func (m *Mission) GetNavigation() *Navigation {
	return m.navigation
}

// GetVersion .
func (m *Mission) GetVersion() Version {
	return m.version
}

// GetNewVersion .
func (m *Mission) GetNewVersion() Version {
	return m.newVersion
}

// NameMission .
func (m *Mission) NameMission(name string) error {
	if m.isCarbonCopy {
		return ErrCannotChange
	}
	m.name = name
	m.newVersion = m.gen.NewVersion()
	return nil
}

// ReplaceNavigationWith .
func (m *Mission) ReplaceNavigationWith(navigation *Navigation) error {
	if m.isCarbonCopy {
		return ErrCannotChange
	}
	m.navigation = navigation
	m.newVersion = m.gen.NewVersion()
	return nil
}

// Generator .
type Generator interface {
	NewID() ID
	NewVersion() Version
}
