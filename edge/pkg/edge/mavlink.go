package edge

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
	mavsdk_rpc_core "edge/pkg/protos/core"
	mavsdk_rpc_mission "edge/pkg/protos/mission"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// TIMEOUT const
const TIMEOUT = 5 * time.Second

// PublishInterval const
const PublishInterval = 500 * time.Millisecond

// Mavlink struct
type Mavlink struct {
	Gr             *grpc.ClientConn
	VehicleID      uint64
	Path           []float64
	Quat           []float64
	Armed          bool
	FlightMode     string
	VideoStreaming bool
}

var rwm sync.RWMutex

// NewMavlink constructor
func NewMavlink(gr *grpc.ClientConn) *Mavlink {
	return &Mavlink{Gr: gr}
}

// SendTelemetry method
func (mavlink *Mavlink) SendTelemetry(cloud string) {
	t := time.NewTicker(PublishInterval)
	for {
		<-t.C

		rwmLocker := rwm.RLocker()
		rwmLocker.Lock()

		if mavlink.VehicleID == 0 {
			log.Printf("continue ...")
			rwmLocker.Unlock()
			continue
		}

		posdata := Telemetry{
			ID:           strconv.FormatUint(mavlink.VehicleID, 10),
			Latitude:     mavlink.Path[1],
			Longitude:    mavlink.Path[0],
			Altitude:     mavlink.Path[2],
			Speed:        0,
			Armed:        mavlink.Armed,
			FlightMode:   mavlink.FlightMode,
			OrientationX: mavlink.Quat[0],
			OrientationY: mavlink.Quat[1],
			OrientationZ: mavlink.Quat[2],
			OrientationW: mavlink.Quat[3],
		}

		jsonData, _ := json.Marshal(posdata)
		log.Printf("Send CLOUD data=%s\n", jsonData)

		req, err := http.NewRequest(
			"POST",
			"http://"+cloud+"/api/v1/communications/"+posdata.ID+"/telemetries",
			strings.NewReader(string(jsonData)),
		)
		if err != nil {
			log.Println("telemetry request error:", err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("telemetry request error:", err)
		}
		defer resp.Body.Close()

		rwmLocker.Unlock()
	}
}

// Listen method
func (mavlink *Mavlink) Listen() {

	core := mavsdk_rpc_core.NewCoreServiceClient(mavlink.Gr)
	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(mavlink.Gr)
	action := mavsdk_rpc_action.NewActionServiceClient(mavlink.Gr)
	mission := mavsdk_rpc_mission.NewMissionServiceClient(mavlink.Gr)

	connStateStream := func(mavlink *Mavlink, core mavsdk_rpc_core.CoreServiceClient) <-chan *mavsdk_rpc_core.ConnectionState {
		connStateStream := make(chan *mavsdk_rpc_core.ConnectionState)

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		connStateRequest := mavsdk_rpc_core.SubscribeConnectionStateRequest{}
		connStateReceiver, err := core.SubscribeConnectionState(ctx, &connStateRequest)
		if err != nil {
			log.Fatal("connState request error:", err)
		}

		go func() {
			defer close(connStateStream)
			defer cancel()
			for {
				response, err := connStateReceiver.Recv()
				if err == io.EOF {
					log.Println("connState response error:", err)
					return
				}
				if err != nil {
					log.Println("connState response error:", err)
					return
				}
				connState := response.GetConnectionState()
				connStateStream <- connState
			}
		}()
		return connStateStream
	}(mavlink, core)

	positionStream := func(mavlink *Mavlink, telemetry mavsdk_rpc_telemetry.TelemetryServiceClient) <-chan *mavsdk_rpc_telemetry.Position {
		positionStream := make(chan *mavsdk_rpc_telemetry.Position)

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		positionRequest := mavsdk_rpc_telemetry.SubscribePositionRequest{}
		positionReceiver, err := telemetry.SubscribePosition(ctx, &positionRequest)
		if err != nil {
			log.Fatal("position request error:", err)
		}

		go func() {
			defer close(positionStream)
			defer cancel()
			for {
				response, err := positionReceiver.Recv()
				if err == io.EOF {
					log.Println("position response error:", err)
					return
				}
				if err != nil {
					log.Println("position response error:", err)
					return
				}
				position := response.GetPosition()
				positionStream <- position
			}
		}()
		return positionStream
	}(mavlink, telemetry)

	quaternionStream := func(mavlink *Mavlink, telemetry mavsdk_rpc_telemetry.TelemetryServiceClient) <-chan *mavsdk_rpc_telemetry.Quaternion {
		quaternionStream := make(chan *mavsdk_rpc_telemetry.Quaternion)

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		quaternionRequest := mavsdk_rpc_telemetry.SubscribeAttitudeQuaternionRequest{}
		quaternionReceiver, err := telemetry.SubscribeAttitudeQuaternion(ctx, &quaternionRequest)
		if err != nil {
			log.Fatal("quaternion request error:", err)
		}

		go func() {
			defer close(quaternionStream)
			defer cancel()
			for {
				response, err := quaternionReceiver.Recv()
				if err == io.EOF {
					log.Println("quaternion response error:", err)
					return
				}
				if err != nil {
					log.Println("quaternion response error:", err)
					return
				}
				quaternion := response.GetAttitudeQuaternion()
				quaternionStream <- quaternion
			}
		}()
		return quaternionStream
	}(mavlink, telemetry)

	armedStream := func(mavlink *Mavlink, telemetry mavsdk_rpc_telemetry.TelemetryServiceClient) <-chan bool {
		armedStream := make(chan bool)

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		armedRequest := mavsdk_rpc_telemetry.SubscribeArmedRequest{}
		armedReceiver, err := telemetry.SubscribeArmed(ctx, &armedRequest)
		if err != nil {
			log.Fatal("armed request error:", err)
		}

		go func() {
			defer close(armedStream)
			defer cancel()
			for {
				response, err := armedReceiver.Recv()
				if err == io.EOF {
					log.Println("armed response error:", err)
					return
				}
				if err != nil {
					log.Println("armed response error:", err)
					return
				}
				armed := response.GetIsArmed()
				armedStream <- armed
			}
		}()
		return armedStream
	}(mavlink, telemetry)

	flightModeStream := func(mavlink *Mavlink, telemetry mavsdk_rpc_telemetry.TelemetryServiceClient) <-chan mavsdk_rpc_telemetry.FlightMode {
		flightModeStream := make(chan mavsdk_rpc_telemetry.FlightMode)

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		flightModeRequest := mavsdk_rpc_telemetry.SubscribeFlightModeRequest{}
		flightModeReceiver, err := telemetry.SubscribeFlightMode(ctx, &flightModeRequest)
		if err != nil {
			log.Fatal("flightMode request error:", err)
		}

		go func() {
			defer close(flightModeStream)
			defer cancel()
			for {
				response, err := flightModeReceiver.Recv()
				if err == io.EOF {
					log.Println("flightMode response error:", err)
					return
				}
				if err != nil {
					log.Println("flightMode response error:", err)
					return
				}
				flightMode := response.GetFlightMode()
				flightModeStream <- flightMode
			}
		}()
		return flightModeStream
	}(mavlink, telemetry)

	go func(mavlink *Mavlink,
		connStateStream <-chan *mavsdk_rpc_core.ConnectionState,
		positionStream <-chan *mavsdk_rpc_telemetry.Position,
		quaternionStream <-chan *mavsdk_rpc_telemetry.Quaternion,
		armedStream <-chan bool,
		flightModeStream <-chan mavsdk_rpc_telemetry.FlightMode) {

		for {
			select {
			case connState := <-connStateStream:
				rwm.Lock()
				mavlink.VehicleID = connState.GetUuid()
				rwm.Unlock()
			case position := <-positionStream:
				rwm.Lock()
				path := []float64{}
				path = append(path, position.GetLongitudeDeg())
				path = append(path, position.GetLatitudeDeg())
				path = append(path, float64(position.GetAbsoluteAltitudeM()))
				path = append(path, float64(position.GetRelativeAltitudeM()))
				mavlink.Path = path
				rwm.Unlock()
			case quaternion := <-quaternionStream:
				rwm.Lock()
				quat := []float64{}
				quat = append(quat, float64(quaternion.GetX()))
				quat = append(quat, float64(quaternion.GetY()))
				quat = append(quat, float64(quaternion.GetZ()))
				quat = append(quat, float64(quaternion.GetW()))
				mavlink.Quat = quat
				rwm.Unlock()
			case armed := <-armedStream:
				rwm.Lock()
				mavlink.Armed = armed
				rwm.Unlock()
			case flightMode := <-flightModeStream:
				rwm.Lock()
				mavlink.FlightMode = flightMode.String()
				rwm.Unlock()
			}
		}
	}(mavlink, connStateStream, positionStream, quaternionStream, armedStream, flightModeStream)

	armCommandStream, disarmCommandStream, takeoffCommandStream, landCommandStream, rtlCommandStream, uploadCommandStream, startCommandStream, pauseCommandStream, streamingCommandStream :=
		func(mavlink *Mavlink) (<-chan *Command, <-chan *Command, <-chan *Command, <-chan *Command, <-chan *Command, <-chan *Command, <-chan *Command, <-chan *Command, <-chan *Command) {

			armCommandStream := make(chan *Command)
			disarmCommandStream := make(chan *Command)
			takeoffCommandStream := make(chan *Command)
			landCommandStream := make(chan *Command)
			rtlCommandStream := make(chan *Command)
			uploadCommandStream := make(chan *Command)
			startCommandStream := make(chan *Command)
			pauseCommandStream := make(chan *Command)
			streamingCommandStream := make(chan *Command)

			// var command Command
			// go func(mavlink *Mavlink) {
			// 	defer close(armCommandStream)
			// 	defer close(disarmCommandStream)
			// 	defer close(takeoffCommandStream)
			// 	defer close(landCommandStream)
			// 	defer close(rtlCommandStream)
			// 	defer close(uploadCommandStream)
			// 	defer close(startCommandStream)
			// 	defer close(pauseCommandStream)
			// 	defer close(streamingCommandStream)
			// 	for {
			// 		websocket.JSON.Receive(mavlink.Ws, &command)

			// 		switch command.MessageID {
			// 		case "arm":
			// 			armCommandStream <- &command
			// 		case "disarm":
			// 			disarmCommandStream <- &command
			// 		case "takeoff":
			// 			takeoffCommandStream <- &command
			// 		case "land":
			// 			landCommandStream <- &command
			// 		case "rtl":
			// 			rtlCommandStream <- &command
			// 		case "upload":
			// 			uploadCommandStream <- &command
			// 		case "start":
			// 			startCommandStream <- &command
			// 		case "pause":
			// 			pauseCommandStream <- &command
			// 		case "streamon":
			// 			streamingCommandStream <- &command
			// 		case "streamoff":
			// 			streamingCommandStream <- &command
			// 		default:
			// 			continue
			// 		}
			// 		log.Printf("Receive CLOUD data=%s, %s, %s\n", command.VehicleID, command.MessageID, string(command.Payload))
			// 	}
			// }(mavlink)

			return armCommandStream, disarmCommandStream, takeoffCommandStream, landCommandStream, rtlCommandStream, uploadCommandStream, startCommandStream, pauseCommandStream, streamingCommandStream
		}(mavlink)

	go func(mavlink *Mavlink,
		armCommandStream <-chan *Command,
		action mavsdk_rpc_action.ActionServiceClient) <-chan *Command {

		for {
			command := <-armCommandStream

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			armRequest := mavsdk_rpc_action.ArmRequest{}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, "")
			response, err := action.Arm(ctx, &armRequest)
			if err != nil {
				log.Println("Receive MAVSDK arm response error:", err)
			}
			log.Println("Receive MAVSDK arm response status:", response.GetActionResult().GetResultStr())
		}
	}(mavlink, armCommandStream, action)

	go func(mavlink *Mavlink,
		disarmCommandStream <-chan *Command,
		action mavsdk_rpc_action.ActionServiceClient) <-chan *Command {

		for {
			command := <-disarmCommandStream

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			disarmRequest := mavsdk_rpc_action.DisarmRequest{}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, "")
			response, err := action.Disarm(ctx, &disarmRequest)
			if err != nil {
				log.Println("Receive MAVSDK disarm response error:", err)
			}
			log.Println("Receive MAVSDK disarm response status:", response.GetActionResult().GetResultStr())
		}
	}(mavlink, disarmCommandStream, action)

	go func(mavlink *Mavlink,
		takeoffCommandStream <-chan *Command,
		action mavsdk_rpc_action.ActionServiceClient) <-chan *Command {

		for {
			command := <-takeoffCommandStream

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			takeoffRequest := mavsdk_rpc_action.TakeoffRequest{}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, "")
			response, err := action.Takeoff(ctx, &takeoffRequest)
			if err != nil {
				log.Println("Receive MAVSDK takeoff response error:", err)
			}
			log.Println("Receive MAVSDK takeoff response status:", response.GetActionResult().GetResultStr())
		}
	}(mavlink, takeoffCommandStream, action)

	go func(mavlink *Mavlink,
		landCommandStream <-chan *Command,
		action mavsdk_rpc_action.ActionServiceClient) <-chan *Command {

		for {
			command := <-landCommandStream

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			landRequest := mavsdk_rpc_action.LandRequest{}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, "")
			response, err := action.Land(ctx, &landRequest)
			if err != nil {
				log.Println("Receive MAVSDK land response error:", err)
			}
			log.Println("Receive MAVSDK land response status:", response.GetActionResult().GetResultStr())
		}
	}(mavlink, landCommandStream, action)

	go func(mavlink *Mavlink,
		rtlCommandStream <-chan *Command,
		action mavsdk_rpc_action.ActionServiceClient) <-chan *Command {

		for {
			command := <-rtlCommandStream

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			rtlRequest := mavsdk_rpc_action.ReturnToLaunchRequest{}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, "")
			response, err := action.ReturnToLaunch(ctx, &rtlRequest)
			if err != nil {
				log.Println("Receive MAVSDK rtl response error:", err)
			}
			log.Println("Receive MAVSDK rtl response status:", response.GetActionResult().GetResultStr())
		}
	}(mavlink, rtlCommandStream, action)

	go func(mavlink *Mavlink,
		uploadCommandStream <-chan *Command,
		mission mavsdk_rpc_mission.MissionServiceClient) <-chan *Command {

		for {
			command := <-uploadCommandStream

			var missionCommand Mission
			json.Unmarshal(command.Payload, &missionCommand)

			missionItems := make([]*mavsdk_rpc_mission.MissionItem, 0)
			for _, each := range missionCommand.MissionItems {
				missionItem := &mavsdk_rpc_mission.MissionItem{
					LatitudeDeg:       each.Lat,
					LongitudeDeg:      each.Lon,
					RelativeAltitudeM: each.Alt,
					SpeedMS:           each.Speed,
				}
				missionItems = append(missionItems, missionItem)
			}

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			uploadRequest := mavsdk_rpc_mission.UploadMissionRequest{
				MissionItems: missionItems,
			}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, uploadRequest.String())
			response, err := mission.UploadMission(ctx, &uploadRequest)
			if err != nil {
				log.Println("Receive MAVSDK upload response error:", err)
			}
			log.Println("Receive MAVSDK upload response status:", response.GetMissionResult().GetResultStr())
		}
	}(mavlink, uploadCommandStream, mission)

	go func(mavlink *Mavlink,
		startCommandStream <-chan *Command,
		mission mavsdk_rpc_mission.MissionServiceClient) <-chan *Command {

		for {
			command := <-startCommandStream

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			startRequest := mavsdk_rpc_mission.StartMissionRequest{}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, "")
			response, err := mission.StartMission(ctx, &startRequest)
			if err != nil {
				log.Println("Receive MAVSDK start response error:", err)
			}
			log.Println("Receive MAVSDK start response status:", response.GetMissionResult().GetResultStr())
		}
	}(mavlink, startCommandStream, mission)

	go func(mavlink *Mavlink,
		pauseCommandStream <-chan *Command,
		mission mavsdk_rpc_mission.MissionServiceClient) <-chan *Command {

		for {
			command := <-pauseCommandStream

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
			defer cancel()

			pauseRequest := mavsdk_rpc_mission.PauseMissionRequest{}
			log.Printf("Send MAVSDK data=%s, %s, %s\n", command.VehicleID, command.MessageID, "")
			response, err := mission.PauseMission(ctx, &pauseRequest)
			if err != nil {
				log.Println("Receive MAVSDK pause response error:", err)
			}
			log.Println("Receive MAVSDK pause response status:", response.GetMissionResult().GetResultStr())
		}
	}(mavlink, pauseCommandStream, mission)

	go func(mavlink *Mavlink,
		streamingCommandStream <-chan *Command) {

		var done chan interface{}
		defer func() {
			if done != nil {
				close(done)
			}
		}()

		for {
			command := <-streamingCommandStream

			if command.MessageID == "streamon" {
				if done == nil {
					done = make(chan interface{})
					// go streaming(done, command.VehicleID)

					rwm.Lock()
					mavlink.VideoStreaming = true
					rwm.Unlock()
				}
			} else if command.MessageID == "streamoff" {
				if done != nil {
					close(done)
					done = nil

					rwm.Lock()
					mavlink.VideoStreaming = false
					rwm.Unlock()
				}
			}
		}
	}(mavlink, streamingCommandStream)

}
