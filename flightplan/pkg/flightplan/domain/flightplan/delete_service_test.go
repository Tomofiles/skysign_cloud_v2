package flightplan

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		generator:   nil,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteFlightplan(ctx, repo, pub, DefaultID)

	expectEvent := DeletedEvent{id: DefaultID}

	a.Len(repo.deleteIDs, 1)
	a.Equal(repo.deleteIDs[0], DefaultID)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

func TestNotFoundFlightplanWhenDeleteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(nil, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteFlightplan(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, errors.New("flightplan not found"))
}

func TestGetErrorWhenDeleteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(nil, errors.New("get error"))
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteFlightplan(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, errors.New("get error"))
}

func TestDeleteErrorWhenDeleteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		generator:   nil,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Delete", mock.Anything).Return(errors.New("delete error"))

	pub := &publisherMock{}

	ret := DeleteFlightplan(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, errors.New("delete error"))
}

type repositoryMockDeleteService struct {
	mock.Mock

	saveFlightplans []*Flightplan
	deleteIDs       []ID
}

func (rm *repositoryMockDeleteService) GetByID(ctx context.Context, id ID) (*Flightplan, error) {
	ret := rm.Called(id)
	var f *Flightplan
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*Flightplan)
	}
	return f, ret.Error(1)
}
func (rm *repositoryMockDeleteService) Save(ctx context.Context, f *Flightplan) error {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) Delete(ctx context.Context, id ID) error {
	ret := rm.Called(id)
	if ret.Error(0) == nil {
		rm.deleteIDs = append(rm.deleteIDs, id)
	}
	return ret.Error(0)
}
