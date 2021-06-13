package vehicle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Vehicleの名前を変更する。
// バージョンが更新されていることを検証する。
func TestChangeVehiclesName(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	vehicle := NewInstance(gen)

	err := vehicle.NameVehicle(DefaultName)

	a.Equal(vehicle.GetName(), DefaultName)
	a.Equal(vehicle.GetVersion(), DefaultVersion1)
	a.Equal(vehicle.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Vehicleの名前を変更する。
// カーボンコピーされたVehicleの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenChangeVehiclesName(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultName1    = DefaultName + "-1"
		DefaultName2    = DefaultName + "-2"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion2},
	}
	original := &Vehicle{
		id:              DefaultID,
		name:            DefaultName1,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion1,
		newVersion:      DefaultVersion1,
	}
	vehicle := Copy(gen, CopiedID, original)

	err := vehicle.NameVehicle(DefaultName2)

	a.Equal(vehicle.GetName(), DefaultName1)
	a.Equal(vehicle.GetVersion(), DefaultVersion1)
	a.Equal(vehicle.GetNewVersion(), DefaultVersion1)
	a.Equal(err, ErrCannotChange)
}

// VehicleにCommunicationIDを付与する。
// CommunicationIDを付与したら、イベントを生成して発行する。
// その際、新しいCommunicationIDのみ、購読者に通知されることを検証する。
// また、バージョンが更新されていることを検証する。
func TestChangeNewVehiclesCommunicationIdAndPublishEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	pub := &publisherMock{}
	vehicle := NewInstance(gen)

	vehicle.SetPublisher(pub)
	err := vehicle.GiveCommunication(DefaultCommunicationID)

	expectEvent := CommunicationIdGaveEvent{
		CommunicationID: DefaultCommunicationID,
	}

	a.Equal(vehicle.GetCommunicationID(), DefaultCommunicationID)
	a.Equal(vehicle.GetVersion(), DefaultVersion1)
	a.Equal(vehicle.GetNewVersion(), DefaultVersion2)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
	a.Nil(err)
}
