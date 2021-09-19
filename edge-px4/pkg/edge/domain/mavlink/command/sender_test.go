package command

import (
	"context"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"

	"github.com/stretchr/testify/assert"
)

// TestCommandSender .
func TestCommandSender(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan struct{})
	fnc := func() error {
		return nil
	}
	name := "test command sender"

	sendExit := CommandSender(ctx, supportMock, stream, fnc, name)

	stream <- struct{}{}
	close(stream)

	<-sendExit

	a.Equal([]string{"command test command sender close"}, supportMock.messages)
}

// TestCommandSenderContextDone .
func TestCommandSenderContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}
	stream := make(chan struct{})
	fnc := func() error {
		return nil
	}
	name := "test command sender"

	sendExit := CommandSender(ctx, supportMock, stream, fnc, name)

	cancel()

	<-sendExit

	a.Equal([]string{"command test command sender done"}, supportMock.messages)
}

// TestErrorWhenCommandSender .
func TestErrorWhenCommandSender(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan struct{})
	fnc := func() error {
		return ErrSend
	}
	name := "test command sender"

	sendExit := CommandSender(ctx, supportMock, stream, fnc, name)

	stream <- struct{}{}
	close(stream)

	<-sendExit

	a.Equal([]string{"command test command sender error: send error", "command test command sender close"}, supportMock.messages)
}

// TestMissionSender .
func TestMissionSender(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *model.Mission)
	var resMission *model.Mission
	fnc := func(mission *model.Mission) error {
		resMission = mission
		return nil
	}

	sendExit := MissionSender(ctx, supportMock, stream, fnc)

	response1 := &model.Mission{}
	stream <- response1
	close(stream)

	<-sendExit

	a.Equal(response1, resMission)
	a.Equal([]string{"mission close"}, supportMock.messages)
}

// TestMissionSenderContextDone .
func TestMissionSenderContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}
	stream := make(chan *model.Mission)
	var resMission *model.Mission
	fnc := func(mission *model.Mission) error {
		resMission = mission
		return nil
	}

	sendExit := MissionSender(ctx, supportMock, stream, fnc)

	cancel()

	<-sendExit

	a.Nil(resMission)
	a.Equal([]string{"mission done"}, supportMock.messages)
}

// TestErrorWhenMissionSender .
func TestErrorWhenMissionSender(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *model.Mission)
	var resMission *model.Mission
	fnc := func(mission *model.Mission) error {
		resMission = mission
		return ErrSend
	}

	sendExit := MissionSender(ctx, supportMock, stream, fnc)

	response1 := &model.Mission{}
	stream <- response1
	close(stream)

	<-sendExit

	a.Equal(response1, resMission)
	a.Equal([]string{"mission error: send error", "mission close"}, supportMock.messages)
}
