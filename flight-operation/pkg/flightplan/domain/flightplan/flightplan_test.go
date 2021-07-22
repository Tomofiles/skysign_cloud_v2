package flightplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Flightplanの名前を変更する。
// バージョンが更新されていることを検証する。
func TestChangeFlightplansName(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	flightplan := NewInstance(gen)

	err := flightplan.NameFlightplan(DefaultName)

	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Flightplanの説明を変更する。
// バージョンが更新されていることを検証する。
func TestChangeFlightplansDescription(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	flightplan := NewInstance(gen)

	err := flightplan.ChangeDescription(DefaultDescription)

	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Flightplanの機体数を変更する。
// バージョンが更新されていることを検証する。
// イベントパブリッシャを設定していないため、イベント発行がスキップされること。
func TestNoEventsWhenChangeFlightplansNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		fleetID:  DefaultFleetID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	flightplan := NewInstance(gen)

	err := flightplan.ChangeNumberOfVehicles(DefaultNumberOfVehicles)

	a.Equal(flightplan.GetFleetID(), DefaultFleetID)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Flightplanの機体数を変更する。
// バージョンが更新されていることを検証する。
// 古いFleetIDが存在しないため、FleetIDが付与されたイベントのみ
// 発行されることを検証する。
func TestOldFleetNotExistsWhenChangeFlightplansNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		fleetID:  DefaultFleetID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	pub := &publisherMock{}
	flightplan := NewInstance(gen)
	flightplan.SetPublisher(pub)

	err := flightplan.ChangeNumberOfVehicles(DefaultNumberOfVehicles)

	expectEvent := FleetIDGaveEvent{
		FleetID:          DefaultFleetID,
		NumberOfVehicles: DefaultNumberOfVehicles,
	}

	a.Equal(flightplan.GetFleetID(), DefaultFleetID)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
	a.Nil(err)
}

// Flightplanの機体数を変更する。
// バージョンが更新されていることを検証する。
// 古いFleetIDが存在するため、FleetIDが付与されたイベントと
// FleetIDが削除されたイベントの両方が発行されることを検証する。
func TestChangeFlightplansNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		NewFleetID      = DefaultFleetID + "-new"
	)

	gen := &generatorMock{
		id:       DefaultID,
		fleetID:  NewFleetID,
		versions: []Version{DefaultVersion2},
	}
	pub := &publisherMock{}
	flightplan := &Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		gen:         gen,
	}
	flightplan.SetPublisher(pub)

	err := flightplan.ChangeNumberOfVehicles(DefaultNumberOfVehicles)

	expectEvent1 := FleetIDRemovedEvent{
		FleetID: DefaultFleetID,
	}
	expectEvent2 := FleetIDGaveEvent{
		FleetID:          NewFleetID,
		NumberOfVehicles: DefaultNumberOfVehicles,
	}

	a.Equal(flightplan.GetFleetID(), NewFleetID)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Len(pub.events, 2)
	a.Equal(pub.events, []interface{}{expectEvent1, expectEvent2})
	a.Nil(err)
}

// FlightplanのFleetIDを削除する。
// バージョンが更新されていることを検証する。
// イベントパブリッシャを設定していないため、イベント発行がスキップされること。
func TestNoPublisherWhenRemoveFlightplansFleetID(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	flightplan := &Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		gen:         gen,
	}

	err := flightplan.RemoveFleetID()

	a.Equal(flightplan.GetFleetID(), BlankFleetID)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// FlightplanのFleetIDを削除する。
// バージョンが更新されていないことを検証する。
// 古いFleetIDが存在しないため、ドメインイベントが発行されないことを検証する。
func TestOldFleetNotExistsWhenRemoveFlightplansFleetID(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		fleetID:  DefaultFleetID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	pub := &publisherMock{}
	flightplan := NewInstance(gen)
	flightplan.SetPublisher(pub)

	err := flightplan.RemoveFleetID()

	a.Equal(flightplan.GetFleetID(), BlankFleetID)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion1)
	a.Len(pub.events, 0)
	a.Nil(err)
}

// FlightplanのFleetIDを削除する。
// バージョンが更新されていることを検証する。
// 古いFleetIDが存在するため、FleetIDが削除されたイベントが
// 発行されることを検証する。
func TestRemoveFlightplansFleetID(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	pub := &publisherMock{}
	flightplan := &Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		gen:         gen,
		pub:         pub,
	}

	err := flightplan.RemoveFleetID()

	expectEvent := FleetIDRemovedEvent{
		FleetID: DefaultFleetID,
	}

	a.Equal(flightplan.GetFleetID(), BlankFleetID)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
	a.Nil(err)
}
