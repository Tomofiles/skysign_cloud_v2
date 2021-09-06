package command

import (
	"context"
	"edge/pkg/edge"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCommandDistributerContextDone .
func TestCommandDistributerContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	cancel()

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.False(resDisarm)
	a.False(resStart)
	a.False(resPause)
	a.False(resTakeoff)
	a.False(resLand)
	a.False(resReturn)
	a.Equal([]string{"command distributer done"}, supportMock.messages)
}

// TestCommandDistributerARM .
func TestCommandDistributerARM(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "ARM",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.True(resArm)
	a.False(resDisarm)
	a.False(resStart)
	a.False(resPause)
	a.False(resTakeoff)
	a.False(resLand)
	a.False(resReturn)
	a.Equal([]string{"command distributer close"}, supportMock.messages)
}

// TestCommandDistributerDISARM .
func TestCommandDistributerDISARM(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "DISARM",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.True(resDisarm)
	a.False(resStart)
	a.False(resPause)
	a.False(resTakeoff)
	a.False(resLand)
	a.False(resReturn)
	a.Equal([]string{"command distributer close"}, supportMock.messages)
}

// TestCommandDistributerSTART .
func TestCommandDistributerSTART(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "START",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.False(resDisarm)
	a.True(resStart)
	a.False(resPause)
	a.False(resTakeoff)
	a.False(resLand)
	a.False(resReturn)
	a.Equal([]string{"command distributer close"}, supportMock.messages)
}

// TestCommandDistributerPAUSE .
func TestCommandDistributerPAUSE(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "PAUSE",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.False(resDisarm)
	a.False(resStart)
	a.True(resPause)
	a.False(resTakeoff)
	a.False(resLand)
	a.False(resReturn)
	a.Equal([]string{"command distributer close"}, supportMock.messages)
}

// TestCommandDistributerTAKEOFF .
func TestCommandDistributerTAKEOFF(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "TAKEOFF",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.False(resDisarm)
	a.False(resStart)
	a.False(resPause)
	a.True(resTakeoff)
	a.False(resLand)
	a.False(resReturn)
	a.Equal([]string{"command distributer close"}, supportMock.messages)
}

// TestCommandDistributerLAND .
func TestCommandDistributerLAND(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "LAND",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.False(resDisarm)
	a.False(resStart)
	a.False(resPause)
	a.False(resTakeoff)
	a.True(resLand)
	a.False(resReturn)
	a.Equal([]string{"command distributer close"}, supportMock.messages)
}

// TestCommandDistributerRETURN .
func TestCommandDistributerRETURN(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "RETURN",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.False(resDisarm)
	a.False(resStart)
	a.False(resPause)
	a.False(resTakeoff)
	a.False(resLand)
	a.True(resReturn)
	a.Equal([]string{"command distributer close"}, supportMock.messages)
}

// TestCommandDistributerNONE .
func TestCommandDistributerNONE(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Command)

	ret := CommandDistributer(ctx, supportMock, stream)

	command := &edge.Command{
		Type: "NONE",
	}
	stream <- command
	close(stream)

	resArm, resDisarm, resStart, resPause, resTakeoff, resLand, resReturn := WaitCommandSender(ret)

	a.False(resArm)
	a.False(resDisarm)
	a.False(resStart)
	a.False(resPause)
	a.False(resTakeoff)
	a.False(resLand)
	a.False(resReturn)
	a.Equal([]string{"command cannot distribute", "command distributer close"}, supportMock.messages)
}

// TestMissionDistributerContextDone .
func TestMissionDistributerContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}
	stream := make(chan *edge.Mission)

	ret := MissionDistributer(ctx, supportMock, stream)

	cancel()

	var wg sync.WaitGroup
	var resMission bool

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-ret:
				if !ok {
					return
				}
				resMission = true
			}
		}
	}()

	wg.Wait()

	a.False(resMission)
	a.Equal([]string{"mission distributer done"}, supportMock.messages)
}

// TestMissionDistributer .
func TestMissionDistributer(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}
	stream := make(chan *edge.Mission)

	ret := MissionDistributer(ctx, supportMock, stream)

	mission := &edge.Mission{}
	stream <- mission
	close(stream)

	var wg sync.WaitGroup
	var resMission bool

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-ret:
				if !ok {
					return
				}
				resMission = true
			}
		}
	}()

	wg.Wait()

	a.True(resMission)
	a.Equal([]string{"mission distributer close"}, supportMock.messages)
}
