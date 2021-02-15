package flightplan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &testGenerator{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	repo := &repositoryMock{}
	pub := &publisherMock{}

	CreateNewFlightplan(ctx, gen, repo, pub, DefaultName, DefaultDescription)

	expectFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion3,
		generator:   gen,
	}
	expectEvent := CreatedEvent{}
	if a.Len(repo.flightplans, 1) {
		a.Equal(repo.flightplans[0], &expectFlightplan)
	}
	if a.Len(pub.events, 1) {
		a.Equal(pub.events[0], expectEvent)
	}
}
