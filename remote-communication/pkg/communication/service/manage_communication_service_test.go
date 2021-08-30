package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCommunicationTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &communicationIDCommandMock{
		ID: string(DefaultCommunicationID),
	}
	ret := service.CreateCommunication(command)

	a.Nil(ret)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateCommunicationOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &communicationIDCommandMock{
		ID: string(DefaultCommunicationID),
	}
	ret := service.createCommunicationOperation(
		nil,
		pub,
		command,
	)

	a.Nil(ret)
}

func TestDeleteCommunicationTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("Delete", DefaultCommunicationID).Return(nil)

	service := &manageCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &communicationIDCommandMock{
		ID: string(DefaultCommunicationID),
	}
	ret := service.DeleteCommunication(command)

	a.Nil(ret)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestDeleteCommunicationOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("Delete", DefaultCommunicationID).Return(nil)
	pub := &publisherMock{}

	service := &manageCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &communicationIDCommandMock{
		ID: string(DefaultCommunicationID),
	}
	ret := service.deleteCommunicationOperation(
		nil,
		pub,
		command,
	)

	a.Nil(ret)
}
