package vehicle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Vehicleを一つ新しく作成し、初期状態を検証する。
func TestCreateNewVehicle(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion},
	}
	vehicle := NewInstance(gen)

	a.Equal(vehicle.GetID(), DefaultID)
	a.Empty(vehicle.GetName())
	a.Empty(vehicle.GetCommunicationID())
	a.Equal(vehicle.isCarbonCopy, Original)
	a.Equal(vehicle.GetVersion(), DefaultVersion)
	a.Equal(vehicle.GetNewVersion(), DefaultVersion)
	a.Equal(vehicle.gen, gen)
	a.Nil(vehicle.pub)
}

// Vehicleのカーボンコピーを作成し、初期状態を検証する。
func TestCopyVehicle(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	id := DefaultID + "-copied"
	original := &Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		pub:             &publisherMock{},
	}
	vehicle := Copy(gen, id, original)

	a.Equal(vehicle.GetID(), id)
	a.Equal(vehicle.GetName(), DefaultName)
	a.Equal(vehicle.GetCommunicationID(), DefaultCommunicationID)
	a.Equal(vehicle.isCarbonCopy, CarbonCopy)
	a.Equal(vehicle.GetVersion(), DefaultVersion)
	a.Equal(vehicle.GetNewVersion(), DefaultVersion)
	a.Equal(vehicle.gen, gen)
	a.Nil(vehicle.pub)
}

// Vehicleを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestVehicleAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &vehicleComponentMock{
		id:              string(DefaultID),
		name:            DefaultName,
		communicationID: string(DefaultCommunicationID),
		isCarbonCopy:    CarbonCopy,
		version:         string(DefaultVersion),
	}
	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion},
	}
	vehicle := AssembleFrom(gen, comp)

	a.Equal(vehicle.GetID(), DefaultID)
	a.Equal(vehicle.GetName(), DefaultName)
	a.Equal(vehicle.GetCommunicationID(), DefaultCommunicationID)
	a.Equal(vehicle.isCarbonCopy, CarbonCopy)
	a.Equal(vehicle.GetVersion(), DefaultVersion)
	a.Equal(vehicle.GetNewVersion(), DefaultVersion)
	a.Equal(vehicle.gen, gen)
	a.Nil(vehicle.pub)
}

// Vehicleを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartVehicle(t *testing.T) {
	a := assert.New(t)

	vehicle := &Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    CarbonCopy,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
	}
	comp := &vehicleComponentMock{}
	TakeApart(
		vehicle,
		func(id, name, communicationID, version string, isCarbonCopy bool) {
			comp.id = id
			comp.name = name
			comp.communicationID = communicationID
			comp.isCarbonCopy = isCarbonCopy
			comp.version = version
		},
	)

	expectComp := &vehicleComponentMock{
		id:              string(DefaultID),
		name:            DefaultName,
		communicationID: string(DefaultCommunicationID),
		isCarbonCopy:    CarbonCopy,
		version:         string(DefaultVersion),
	}
	a.Equal(comp, expectComp)
}
