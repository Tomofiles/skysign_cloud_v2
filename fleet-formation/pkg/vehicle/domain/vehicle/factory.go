package vehicle

// NewInstance .
func NewInstance(gen Generator) *Vehicle {
	id := gen.NewID()
	version := gen.NewVersion()
	return &Vehicle{
		id:           id,
		isCarbonCopy: Original,
		version:      version,
		newVersion:   version,
		gen:          gen,
	}
}

// Copy .
func Copy(gen Generator, id ID, original *Vehicle) *Vehicle {
	return &Vehicle{
		id:              id,
		name:            original.name,
		communicationID: original.communicationID,
		isCarbonCopy:    CarbonCopy,
		version:         original.version,
		newVersion:      original.newVersion,
		gen:             gen,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Vehicle {
	return &Vehicle{
		id:              ID(comp.GetID()),
		name:            comp.GetName(),
		communicationID: CommunicationID(comp.GetCommunicationID()),
		isCarbonCopy:    comp.GetIsCarbonCopy(),
		version:         Version(comp.GetVersion()),
		newVersion:      Version(comp.GetVersion()),
		gen:             gen,
	}
}

// TakeApart .
func TakeApart(
	vehicle *Vehicle,
	component func(id, name, communicationID, version string, isCarbonCopy bool),
) {
	component(
		string(vehicle.id),
		vehicle.name,
		string(vehicle.communicationID),
		string(vehicle.version),
		vehicle.isCarbonCopy,
	)
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetCommunicationID() string
	GetIsCarbonCopy() bool
	GetVersion() string
}
