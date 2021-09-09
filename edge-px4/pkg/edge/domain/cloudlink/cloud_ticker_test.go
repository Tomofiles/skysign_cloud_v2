package cloudlink

import (
	"context"
	"edge-px4/pkg/edge/domain/model"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCloudTicker .
func TestCloudTicker(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	mission := &model.Mission{ID: DefaultEdgeMissionID, Waypoints: []*model.Waypoints{}}

	supportMock := &supportMock{}
	tickerMock := &tickerMock{}

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{DefaultEdgeCommandID}}, nil
	}
	pullCommand := func(vehicleID, commandID string) (*model.Command, error) {
		return &model.Command{Type: "UPLOAD"}, nil
	}
	pullUploadMission := func(vehicleID, commandID string) (*model.UploadMission, error) {
		return &model.UploadMission{ID: DefaultEdgeCommandID, MissionID: DefaultEdgeMissionID}, nil
	}
	getUploadMission := func(missionID string) (*model.Mission, error) {
		return mission, nil
	}

	streams := CloudTicker(
		ctx,
		supportMock,
		tickerMock,
		pushTelemetry,
		pullCommand,
		pullUploadMission,
		getUploadMission,
	)

	var wg sync.WaitGroup
	var resMission *model.Mission

	wg.Add(1)
	go func() {
		for {
			select {
			case _, ok := <-streams.CommandStream:
				if !ok {
					continue
				}
			case mission, ok := <-streams.MissionStream:
				if !ok {
					continue
				}
				resMission = mission
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()

	a.Equal(mission, resMission)
	a.Empty(supportMock.messages)
}

// TestInternalErrorWhenCloudTicker .
func TestInternalErrorWhenCloudTicker(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}
	tickerMock := &tickerMock{}

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		cancel() // tickerが動いた後cancelしたいため、無理やりだがここでcancel呼び出し
		return "", nil, ErrPushTelemetry
	}

	streams := CloudTicker(
		ctx,
		supportMock,
		tickerMock,
		pushTelemetry,
		nil,
		nil,
		nil,
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-streams.CommandStream:
				if !ok {
					return
				}
			case _, ok := <-streams.MissionStream:
				if !ok {
					return
				}
			}
		}
	}()
	wg.Wait()

	a.Equal([]string{"cloud ticker error: push telemetry error", "cloud ticker done"}, supportMock.messages)
}

// TestCloudTickerContextDone .
func TestCloudTickerContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}
	tickerMock := &tickerMock{}

	cancel()

	streams := CloudTicker(
		ctx,
		supportMock,
		tickerMock,
		nil,
		nil,
		nil,
		nil)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-streams.CommandStream:
				if !ok {
					return
				}
			case _, ok := <-streams.MissionStream:
				if !ok {
					return
				}
			}
		}
	}()
	wg.Wait()

	a.True(tickerMock.isStop)
	a.Equal([]string{"cloud ticker done"}, supportMock.messages)
}

// TestCloudTickerInternalMission .
func TestCloudTickerInternalMission(t *testing.T) {
	a := assert.New(t)

	mission := &model.Mission{ID: DefaultEdgeMissionID, Waypoints: []*model.Waypoints{}}

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{DefaultEdgeCommandID}}, nil
	}
	pullCommand := func(vehicleID, commandID string) (*model.Command, error) {
		return &model.Command{Type: "UPLOAD"}, nil
	}
	pullUploadMission := func(vehicleID, commandID string) (*model.UploadMission, error) {
		return &model.UploadMission{ID: DefaultEdgeCommandID, MissionID: DefaultEdgeMissionID}, nil
	}
	getUploadMission := func(missionID string) (*model.Mission, error) {
		return mission, nil
	}

	commandStream := make(chan *model.Command, 1)
	missionStream := make(chan *model.Mission, 1)

	err := CloudTickerInternal(
		pushTelemetry,
		pullCommand,
		pullUploadMission,
		getUploadMission,
		commandStream,
		missionStream,
	)

	var wg sync.WaitGroup
	var resMission *model.Mission

	wg.Add(1)
	go func() {
		for {
			select {
			case _, ok := <-commandStream:
				if !ok {
					continue
				}
			case mission, ok := <-missionStream:
				if !ok {
					continue
				}
				resMission = mission
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()

	a.Nil(err)
	a.Equal(mission, resMission)
}

// TestNoCommandIDsCloudTickerInternalCommand .
func TestNoCommandIDsCloudTickerInternalCommand(t *testing.T) {
	a := assert.New(t)

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{}}, nil
	}

	err := CloudTickerInternal(
		pushTelemetry,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	a.Nil(err)
}

// TestSingleCommandIDCloudTickerInternalCommand .
func TestSingleCommandIDCloudTickerInternalCommand(t *testing.T) {
	a := assert.New(t)

	command := &model.Command{Type: "ARM"}

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{DefaultEdgeCommandID}}, nil
	}
	pullCommand := func(vehicleID, commandID string) (*model.Command, error) {
		return command, nil
	}

	commandStream := make(chan *model.Command, 1)
	missionStream := make(chan *model.Mission, 1)

	err := CloudTickerInternal(
		pushTelemetry,
		pullCommand,
		nil,
		nil,
		commandStream,
		missionStream,
	)

	var wg sync.WaitGroup
	var resCommand *model.Command

	wg.Add(1)
	go func() {
		for {
			select {
			case command, ok := <-commandStream:
				if !ok {
					continue
				}
				resCommand = command
				wg.Done()
				return
			case _, ok := <-missionStream:
				if !ok {
					continue
				}
			}
		}
	}()
	wg.Wait()

	a.Nil(err)
	a.Equal(command, resCommand)
}

// TestMultipleCommandIDsCloudTickerInternalCommand .
func TestMultipleCommandIDsCloudTickerInternalCommand(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultEdgeCommandID1 = DefaultEdgeCommandID + "-1"
		DefaultEdgeCommandID2 = DefaultEdgeCommandID + "-2"
		DefaultEdgeCommandID3 = DefaultEdgeCommandID + "-3"
	)

	command1 := &model.Command{Type: "ARM"}
	command2 := &model.Command{Type: "DISARM"}
	command3 := &model.Command{Type: "LAND"}
	commands := []*model.Command{command1, command2, command3}

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{DefaultEdgeCommandID1, DefaultEdgeCommandID2, DefaultEdgeCommandID3}}, nil
	}
	pullCommand := func(vehicleID, commandID string) (*model.Command, error) {
		switch commandID {
		case DefaultEdgeCommandID1:
			return command1, nil
		case DefaultEdgeCommandID2:
			return command2, nil
		case DefaultEdgeCommandID3:
			return command3, nil
		}
		return command1, nil
	}

	commandStream := make(chan *model.Command, 3)
	missionStream := make(chan *model.Mission, 3)

	err := CloudTickerInternal(
		pushTelemetry,
		pullCommand,
		nil,
		nil,
		commandStream,
		missionStream,
	)

	var wg sync.WaitGroup
	var resCommands []*model.Command

	wg.Add(3)
	go func() {
		for {
			select {
			case command, ok := <-commandStream:
				if !ok {
					continue
				}
				resCommands = append(resCommands, command)
				wg.Done()
				if len(resCommands) == 3 {
					return
				}
			case _, ok := <-missionStream:
				if !ok {
					continue
				}
			}
		}
	}()
	wg.Wait()

	a.Nil(err)
	a.Equal(commands, resCommands)
}

// TestPushTelemetryErrorWhenCloudTickerInternal .
func TestPushTelemetryErrorWhenCloudTickerInternal(t *testing.T) {
	a := assert.New(t)

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return "", nil, ErrPushTelemetry
	}

	err := CloudTickerInternal(
		pushTelemetry,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	a.Equal(ErrPushTelemetry, err)
}

// TestPullCommandErrorWhenCloudTickerInternal .
func TestPullCommandErrorWhenCloudTickerInternal(t *testing.T) {
	a := assert.New(t)

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{DefaultEdgeCommandID}}, nil
	}
	pullCommand := func(vehicleID, commandID string) (*model.Command, error) {
		return nil, ErrPullCommand
	}

	err := CloudTickerInternal(
		pushTelemetry,
		pullCommand,
		nil,
		nil,
		nil,
		nil,
	)

	a.Equal(ErrPullCommand, err)
}

// TestPullUploadMissionErrorWhenCloudTickerInternal .
func TestPullUploadMissionErrorWhenCloudTickerInternal(t *testing.T) {
	a := assert.New(t)

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{DefaultEdgeCommandID}}, nil
	}
	pullCommand := func(vehicleID, commandID string) (*model.Command, error) {
		return &model.Command{Type: "UPLOAD"}, nil
	}
	pullUploadMission := func(vehicleID, commandID string) (*model.UploadMission, error) {
		return nil, ErrPullUploadMission
	}

	err := CloudTickerInternal(
		pushTelemetry,
		pullCommand,
		pullUploadMission,
		nil,
		nil,
		nil,
	)

	a.Equal(ErrPullUploadMission, err)
}

// TestGetUploadMissionErrorWhenCloudTickerInternal .
func TestGetUploadMissionErrorWhenCloudTickerInternal(t *testing.T) {
	a := assert.New(t)

	pushTelemetry := func() (string, *model.CommandIDs, error) {
		return DefaultEdgeVehicleID, &model.CommandIDs{CommandIds: []string{DefaultEdgeCommandID}}, nil
	}
	pullCommand := func(vehicleID, commandID string) (*model.Command, error) {
		return &model.Command{Type: "UPLOAD"}, nil
	}
	pullUploadMission := func(vehicleID, commandID string) (*model.UploadMission, error) {
		return &model.UploadMission{ID: DefaultEdgeCommandID, MissionID: DefaultEdgeMissionID}, nil
	}
	getUploadMission := func(missionID string) (*model.Mission, error) {
		return nil, ErrGetUploadMission
	}

	err := CloudTickerInternal(
		pushTelemetry,
		pullCommand,
		pullUploadMission,
		getUploadMission,
		nil,
		nil,
	)

	a.Equal(ErrGetUploadMission, err)
}
