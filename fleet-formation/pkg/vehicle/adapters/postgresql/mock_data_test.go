package postgresql

import "github.com/google/uuid"

var NewVehicleID = func() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
var NewVehicleVersion = func() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
var NewVehicleCommunicationID = func() string {
	id, _ := uuid.NewRandom()
	return id.String()
}

const DefaultVehicleName = "vehicle-name"

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
