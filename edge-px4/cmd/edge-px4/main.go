package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"os/signal"
	"sync"
	"time"

	glog_adapter "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/glog"
	grpc_adapter "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/grpc"
	time_adapter "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/time"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/builder"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/cloudlink"
	mavlink_command "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/mavlink/command"
	mavlink_telemetry "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/mavlink/telemetry"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"
)

var (
	mavsdk = "localhost:50051"
	cloud  = "http://localhost:8080"
)

func run() error {
	flag.Parse()
	defer glog_adapter.Flush()

	mavsdkAddressEnv := os.Getenv("MAVSDK_ADDRESS")
	cloudAddressEnv := os.Getenv("CLOUD_ADDRESS")

	if mavsdkAddressEnv != "" {
		mavsdk = mavsdkAddressEnv
	}
	if cloudAddressEnv != "" {
		cloud = cloudAddressEnv
	}

	support := glog_adapter.NewSupport()
	ticker := time_adapter.NewTicker(500 * time.Millisecond)

	support.NotifyInfo("edge PX4 start")

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gr, err := grpc_adapter.NewGrpcClientConnectionWithBlockAndTimeout(mavsdk)
	if err != nil {
		support.NotifyError("grpc client connection error: %v", err)
		return err
	}

	telemetryStream, err := builder.MavlinkTelemetry(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink telemetry error: %v", err)
		return err
	}

	telemetry := model.NewTelemetry()

	connectionStateExit := mavlink_telemetry.ConnectionStateUpdater(ctx, support, telemetry, telemetryStream.ConnectionStateStream)
	positionExit := mavlink_telemetry.PositionUpdater(ctx, support, telemetry, telemetryStream.PositionStream)
	quaternionExit := mavlink_telemetry.QuaternionUpdater(ctx, support, telemetry, telemetryStream.QuaternionStream)
	velocityExit := mavlink_telemetry.VelocityUpdater(ctx, support, telemetry, telemetryStream.VelocityStream)
	armedExit := mavlink_telemetry.ArmedUpdater(ctx, support, telemetry, telemetryStream.ArmedStream)
	flightModeExit := mavlink_telemetry.FlightModeUpdater(ctx, support, telemetry, telemetryStream.FlightModeStream)

	cloudlinkAdapters := builder.Cloudlink(
		ctx,
		cloud,
		support,
		telemetry,
	)

	commandStream := cloudlink.CloudTicker(
		ctx,
		support,
		ticker,
		cloudlinkAdapters.PushTelemetry,
		cloudlinkAdapters.PullCommand,
		cloudlinkAdapters.PullUploadMission,
		cloudlinkAdapters.GetUploadMission,
	)

	cStream := mavlink_command.CommandDistributer(ctx, support, commandStream.CommandStream)
	mStream := mavlink_command.MissionDistributer(ctx, support, commandStream.MissionStream)

	mavlinkAdapters := builder.MavlinkCommand(ctx, gr)

	armSendExit := mavlink_command.CommandSender(ctx, support, cStream.ArmStream, mavlinkAdapters.AdapterArm, "ARM")
	disarmSendExit := mavlink_command.CommandSender(ctx, support, cStream.DisarmStream, mavlinkAdapters.AdapterDisarm, "DISARM")
	startSendExit := mavlink_command.CommandSender(ctx, support, cStream.StartStream, mavlinkAdapters.AdapterStart, "START")
	pauseSendExit := mavlink_command.CommandSender(ctx, support, cStream.PauseStream, mavlinkAdapters.AdapterPause, "PAUSE")
	takeoffSendExit := mavlink_command.CommandSender(ctx, support, cStream.TakeoffStream, mavlinkAdapters.AdapterTakeoff, "TAKEOFF")
	landSendExit := mavlink_command.CommandSender(ctx, support, cStream.LandStream, mavlinkAdapters.AdapterLand, "LAND")
	returnSendExit := mavlink_command.CommandSender(ctx, support, cStream.ReturnStream, mavlinkAdapters.AdapterReturn, "RETURN")
	uploadSendExit := mavlink_command.MissionSender(ctx, support, mStream, mavlinkAdapters.AdapterUpload)

	exitWaiter := func() <-chan struct{} {
		exits := []<-chan struct{}{
			connectionStateExit,
			positionExit,
			quaternionExit,
			velocityExit,
			armedExit,
			flightModeExit,
			armSendExit,
			disarmSendExit,
			startSendExit,
			pauseSendExit,
			takeoffSendExit,
			landSendExit,
			returnSendExit,
			uploadSendExit,
		}

		stream := make(chan struct{})

		go func() {
			defer close(stream)
			var wg sync.WaitGroup
			for _, exit := range exits {
				wg.Add(1)
				go func(exit <-chan struct{}) {
					<-exit
					wg.Done()
				}(exit)
			}
			wg.Wait()
		}()

		return stream
	}()

	return func() (err error) {
		defer func() {
			cancel()
			<-exitWaiter
			if err != nil {
				support.NotifyError("edge PX4 error: %v", err)
				return
			}
			support.NotifyInfo("edge PX4 end")
		}()
		for {
			select {
			case <-stop:
				err = nil
				return
			case <-connectionStateExit:
				err = errors.New("connectionStateExit")
				return
			case <-positionExit:
				err = errors.New("positionExit")
				return
			case <-quaternionExit:
				err = errors.New("quaternionExit")
				return
			case <-velocityExit:
				err = errors.New("velocityExit")
				return
			case <-armedExit:
				err = errors.New("armedExit")
				return
			case <-flightModeExit:
				err = errors.New("flightModeExit")
				return
			case <-armSendExit:
				err = errors.New("armSendExit")
				return
			case <-disarmSendExit:
				err = errors.New("disarmSendExit")
				return
			case <-startSendExit:
				err = errors.New("startSendExit")
				return
			case <-pauseSendExit:
				err = errors.New("pauseSendExit")
				return
			case <-takeoffSendExit:
				err = errors.New("takeoffSendExit")
				return
			case <-landSendExit:
				err = errors.New("landSendExit")
				return
			case <-returnSendExit:
				err = errors.New("returnSendExit")
				return
			case <-uploadSendExit:
				err = errors.New("uploadSendExit")
				return
			}
		}
	}()
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
