package net.tomofiles.skysign.communication.api;

import java.util.List;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.usecase.CommunicateVehicleService;
import net.tomofiles.skysign.communication.usecase.dto.ControlCommandDto;
import net.tomofiles.skysign.communication.usecase.dto.TelemetryDto;
import proto.skysign.PullCommandRequest;
import proto.skysign.PullCommandResponse;
import proto.skysign.PushTelemetryRequest;
import proto.skysign.PushTelemetryResponse;
import proto.skysign.CommunicationVehicleServiceGrpc.CommunicationVehicleServiceImplBase;

@GRpcService
@Controller
public class CommunicateVehicleEndpoint extends CommunicationVehicleServiceImplBase {

    @Autowired
    private CommunicateVehicleService service;

    @Override
    public void pushTelemetry(PushTelemetryRequest request, StreamObserver<PushTelemetryResponse> responseObserver) {
        TelemetryDto telemetry = new TelemetryDto();
        telemetry.setLatitude(request.getLatitude());
        telemetry.setLongitude(request.getLongitude());
        telemetry.setAltitude(request.getAltitude());
        telemetry.setSpeed(request.getSpeed());
        telemetry.setArmed(request.getArmed());
        telemetry.setFlightMode(request.getFlightMode());
        telemetry.setOrientationX(request.getOrientationX());
        telemetry.setOrientationY(request.getOrientationY());
        telemetry.setOrientationZ(request.getOrientationZ());
        telemetry.setOrientationW(request.getOrientationW());

        List<String> commandIds = this.service.pushTelemetry(request.getId(), telemetry);

        PushTelemetryResponse r = PushTelemetryResponse.newBuilder().setId(request.getId()).addAllCommIds(commandIds).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void pullCommand(PullCommandRequest request, StreamObserver<PullCommandResponse> responseObserver) {
        ControlCommandDto command = this.service.pullCommand(request.getCommId(), request.getId());
        System.out.println(command.getMissionId());

        PullCommandResponse r = PullCommandResponse.newBuilder().setCommId(request.getCommId()).setId(request.getId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}