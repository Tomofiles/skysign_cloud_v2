package flightplan

const DefaultID = ID("flightplan-id")
const DefaultVersion1 = Version("version-1")
const DefaultVersion2 = Version("version-2")
const DefaultVersion3 = Version("version-3")
const DefaultName = "flightplan-name"
const DefaultDescription = "flightplan-description"

type testGenerator struct {
	Generator
	id           ID
	versions     []Version
	versionIndex int
}

func (gen *testGenerator) NewID() ID {
	return gen.id
}
func (gen *testGenerator) NewVersion() Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
}

type publisherMock struct {
	events []interface{}
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}
