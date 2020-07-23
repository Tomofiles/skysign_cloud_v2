package net.tomofiles.skysign.communication.api;

import java.util.List;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.usecase.CommunicateEdgeService;
import net.tomofiles.skysign.communication.usecase.dto.ControlCommandDto;
import net.tomofiles.skysign.communication.usecase.dto.TelemetryDto;
import proto.skysign.CommandType;
import proto.skysign.PullCommandRequest;
import proto.skysign.PullCommandResponse;
import proto.skysign.PushTelemetryRequest;
import proto.skysign.PushTelemetryResponse;
import proto.skysign.CommunicationEdgeServiceGrpc.CommunicationEdgeServiceImplBase;

@GRpcService
@Controller
public class CommunicateEdgeEndpoint extends CommunicationEdgeServiceImplBase {

    @Autowired
    private CommunicateEdgeService service;

    @Override
    public void pushTelemetry(PushTelemetryRequest request, StreamObserver<PushTelemetryResponse> responseObserver) {
        TelemetryDto telemetry = new TelemetryDto();
        telemetry.setLatitude(request.getLatitude());
        telemetry.setLongitude(request.getLongitude());
        telemetry.setAltitude(request.getAltitude());
        telemetry.setRelativeAltitude(request.getRelativeAltitude());
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
        ControlCommandDto command = this.service.pullCommand(request.getId(), request.getCommandId());

        CommandType type = CommandType.valueOf(command.getType().toString());
        PullCommandResponse r = PullCommandResponse.newBuilder().setId(request.getId()).setCommandId(request.getCommandId()).setType(type).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}