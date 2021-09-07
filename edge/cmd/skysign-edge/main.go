package main

import (
	"context"
	glog_adapter "edge/pkg/edge/adapters/glog"
	grpc_adapter "edge/pkg/edge/adapters/grpc"
	time_adapter "edge/pkg/edge/adapters/time"
	"edge/pkg/edge/builder"
	"edge/pkg/edge/domain/cloudlink"
	mavlink_command "edge/pkg/edge/domain/mavlink/command"
	mavlink_telemetry "edge/pkg/edge/domain/mavlink/telemetry"
	"os"
	"os/signal"
	"time"
)

var (
	mavsdk = "localhost:50051"
	cloud  = "http://localhost:8080"
)

func main() {
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

	go func() {
		for {
			t := time.NewTimer(1 * time.Second)
			select {
			case <-t.C:
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)

				gr, err := grpc_adapter.NewGrpcClientConnectionWithBlock(mavsdk)
				if err != nil {
					support.NotifyError("grpc client connection error: %v", err)
					continue
				}

				telemetryStream, err := builder.MavlinkTelemetry(ctx, gr, support)
				if err != nil {
					support.NotifyError("mavlink telemetry error: %v", err)
					cancel()
					continue
				}

				telemetry := mavlink_telemetry.NewTelemetry()
				updateExit := mavlink_telemetry.Updater(
					ctx,
					support,
					telemetry,
					telemetryStream.ConnectionStateStream,
					telemetryStream.PositionStream,
					telemetryStream.QuaternionStream,
					telemetryStream.VelocityStream,
					telemetryStream.ArmedStream,
					telemetryStream.FlightModeStream,
				)

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

				mavlinkAdapters := builder.MavlinkCommand(
					ctx,
					gr,
					support,
				)

				armSendExit := mavlink_command.CommandSender(ctx, support, cStream.ArmStream, mavlinkAdapters.AdapterArm, "ARM")
				disarmSendExit := mavlink_command.CommandSender(ctx, support, cStream.DisarmStream, mavlinkAdapters.AdapterDisarm, "DISARM")
				startSendExit := mavlink_command.CommandSender(ctx, support, cStream.StartStream, mavlinkAdapters.AdapterStart, "START")
				pauseSendExit := mavlink_command.CommandSender(ctx, support, cStream.PauseStream, mavlinkAdapters.AdapterPause, "PAUSE")
				takeoffSendExit := mavlink_command.CommandSender(ctx, support, cStream.TakeoffStream, mavlinkAdapters.AdapterTakeoff, "TAKEOFF")
				landSendExit := mavlink_command.CommandSender(ctx, support, cStream.LandStream, mavlinkAdapters.AdapterLand, "LAND")
				returnSendExit := mavlink_command.CommandSender(ctx, support, cStream.ReturnStream, mavlinkAdapters.AdapterReturn, "RETURN")
				uploadSendExit := mavlink_command.MissionSender(ctx, support, mStream, mavlinkAdapters.AdapterUpload)

				// // 障害時動作確認用
				// go func() {
				// 	t := time.NewTimer(5 * time.Second)
				// 	select {
				// 	case <-t.C:
				// 		cancel()
				// 	}
				// }()

				func() {
					defer func() {
						support.NotifyInfo("main loop exit")
						cancel()
					}()
					for {
						select {
						case <-updateExit:
							return
						case <-armSendExit:
							return
						case <-disarmSendExit:
							return
						case <-startSendExit:
							return
						case <-pauseSendExit:
							return
						case <-takeoffSendExit:
							return
						case <-landSendExit:
							return
						case <-returnSendExit:
							return
						case <-uploadSendExit:
							return
						}
					}
				}()
			}
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	<-stop

	time.Sleep(1 * time.Second)

	defer support.NotifyInfo("Skysign Edge end")
}
