package main

import (
	"context"
	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/adapters/grpc"
	"edge/pkg/edge/builder"
	"edge/pkg/edge/domain/command"
	"edge/pkg/edge/domain/telemetry"
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

	support := glog.NewSupport()

	go func() {
		for {
			t := time.NewTimer(1 * time.Second)
			select {
			case <-t.C:
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)

				gr, err := grpc.NewGrpcClientConnectionWithBlock(mavsdk)
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

				tlm := telemetry.NewTelemetry()
				updateExit := telemetry.Updater(
					ctx,
					support,
					tlm,
					telemetryStream.ConnectionStateStream,
					telemetryStream.PositionStream,
					telemetryStream.QuaternionStream,
					telemetryStream.VelocityStream,
					telemetryStream.ArmedStream,
					telemetryStream.FlightModeStream,
				)

				commandStream := builder.Cloudlink(ctx, cloud, tlm)

				cStream := command.CommandDistributer(ctx, support, commandStream.CommandStream)
				mStream := command.MissionDistributer(ctx, support, commandStream.MissionStream)

				adapters := builder.MavlinkCommand(
					ctx,
					gr,
					support,
				)

				armSendExit := command.CommandSender(ctx, support, cStream.ArmStream, adapters.AdapterArm, "ARM")
				disarmSendExit := command.CommandSender(ctx, support, cStream.DisarmStream, adapters.AdapterDisarm, "DISARM")
				startSendExit := command.CommandSender(ctx, support, cStream.StartStream, adapters.AdapterStart, "START")
				pauseSendExit := command.CommandSender(ctx, support, cStream.PauseStream, adapters.AdapterPause, "PAUSE")
				takeoffSendExit := command.CommandSender(ctx, support, cStream.TakeoffStream, adapters.AdapterTakeoff, "TAKEOFF")
				landSendExit := command.CommandSender(ctx, support, cStream.LandStream, adapters.AdapterLand, "LAND")
				returnSendExit := command.CommandSender(ctx, support, cStream.ReturnStream, adapters.AdapterReturn, "RETURN")
				uploadSendExit := command.MissionSender(ctx, support, mStream, adapters.AdapterUpload)

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

	defer support.NotifyInfo("Skysign Edge end.")
}
