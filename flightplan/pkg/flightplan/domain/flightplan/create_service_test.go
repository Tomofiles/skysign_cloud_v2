package flightplan

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &testGenerator{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := CreateNewFlightplan(ctx, gen, repo, pub, DefaultName, DefaultDescription)

	expectFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion3,
		generator:   gen,
	}
	expectEvent := CreatedEvent{id: DefaultID}

	a.Len(repo.saveFlightplans, 1)
	a.Equal(repo.saveFlightplans[0], &expectFlightplan)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

func TestSaveErrorWhenCreateNewFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &testGenerator{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(errors.New("save error"))
	pub := &publisherMock{}

	ret := CreateNewFlightplan(ctx, gen, repo, pub, DefaultName, DefaultDescription)

	a.Len(repo.saveFlightplans, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, errors.New("save error"))
}

type repositoryMockCreateService struct {
	mock.Mock

	saveFlightplans []*Flightplan
}

func (rm *repositoryMockCreateService) GetByID(ctx context.Context, id ID) (*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) Save(ctx context.Context, f *Flightplan) error {
	ret := rm.Called(f)
	if ret.Error(0) == nil {
		rm.saveFlightplans = append(rm.saveFlightplans, f)
	}
	return ret.Error(0)
}
func (rm *repositoryMockCreateService) Delete(ctx context.Context, id ID) error {
	panic("implement me")
}
