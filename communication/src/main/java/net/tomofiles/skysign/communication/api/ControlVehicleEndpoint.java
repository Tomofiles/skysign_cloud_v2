package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.usecase.ControlVehicleService;
import net.tomofiles.skysign.communication.usecase.dto.ControlCommandType;
import net.tomofiles.skysign.communication.usecase.dto.TelemetryDto;
import proto.skysign.PullTelemetryRequest;
import proto.skysign.PullTelemetryResponse;
import proto.skysign.PushCommandRequest;
import proto.skysign.PushCommandResponse;
import proto.skysign.StandByRequest;
import proto.skysign.StandByResponse;
import proto.skysign.CommunicationUserServiceGrpc.CommunicationUserServiceImplBase;

@GRpcService
@Controller
public class ControlVehicleEndpoint extends CommunicationUserServiceImplBase {
    
    @Autowired
    private ControlVehicleService service;

    @Override
    public void standBy(StandByRequest request, StreamObserver<StandByResponse> responseObserver) {
        this.service.standBy(request.getId(), request.getMissionId());

        StandByResponse r = StandByResponse.newBuilder().setId(request.getId()).setMissionId(request.getMissionId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void pullTelemetry(PullTelemetryRequest request, StreamObserver<PullTelemetryResponse> responseObserver) {
        TelemetryDto telemetry = this.service.pullTelemetry(request.getId());

        PullTelemetryResponse r = PullTelemetryResponse.newBuilder()
                .setLatitude(telemetry.getLatitude())
                .setLongitude(telemetry.getLongitude())
                .setAltitude(telemetry.getAltitude())
                .setSpeed(telemetry.getSpeed())
                .setArmed(telemetry.isArmed())
                .setFlightMode(telemetry.getFlightMode())
                .setOrientationX(telemetry.getOrientationX())
                .setOrientationY(telemetry.getOrientationY())
                .setOrientationZ(telemetry.getOrientationZ())
                .setOrientationW(telemetry.getOrientationW())
                .build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void pushCommand(PushCommandRequest request, StreamObserver<PushCommandResponse> responseObserver) {
        this.service.pushCommand(request.getId(), ControlCommandType.valueOf(request.getType().name()));

        PushCommandResponse r = PushCommandResponse.newBuilder().setId(request.getId()).setType(request.getType()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}