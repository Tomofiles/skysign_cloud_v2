package postgresql

import (
	v "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/domain/vehicle"
)

const DefaultVehicleID = v.ID("vehicle-id")
const DefaultVehicleVersion = v.Version("version")
const DefaultVehicleName = "vehicle-name"
const DefaultVehicleCommunicationID = v.CommunicationID("communication-id")
const DefaultFleetID = v.FleetID("fleet-id")

type vehicleComponentMock struct {
	ID              string
	Name            string
	CommunicationID string
	IsCarbonCopy    bool
	Version         string
}

func (v *vehicleComponentMock) GetID() string {
	return v.ID
}

func (v *vehicleComponentMock) GetName() string {
	return v.Name
}

func (v *vehicleComponentMock) GetCommunicationID() string {
	return v.CommunicationID
}

func (v *vehicleComponentMock) GetIsCarbonCopy() bool {
	return v.IsCarbonCopy
}

func (v *vehicleComponentMock) GetVersion() string {
	return v.Version
}
