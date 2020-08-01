package net.tomofiles.skysign.communication.api;

import java.util.List;
import java.util.NoSuchElementException;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
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
        telemetry.setLatitude(request.getTelemetry().getLatitude());
        telemetry.setLongitude(request.getTelemetry().getLongitude());
        telemetry.setAltitude(request.getTelemetry().getAltitude());
        telemetry.setRelativeAltitude(request.getTelemetry().getRelativeAltitude());
        telemetry.setSpeed(request.getTelemetry().getSpeed());
        telemetry.setArmed(request.getTelemetry().getArmed());
        telemetry.setFlightMode(request.getTelemetry().getFlightMode());
        telemetry.setOrientationX(request.getTelemetry().getOrientationX());
        telemetry.setOrientationY(request.getTelemetry().getOrientationY());
        telemetry.setOrientationZ(request.getTelemetry().getOrientationZ());
        telemetry.setOrientationW(request.getTelemetry().getOrientationW());

        List<String> commandIds;
        try {
            commandIds = this.service.pushTelemetry(request.getId(), telemetry);
        } catch (NoSuchElementException e) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(e)
                    .withDescription(e.getMessage())
                    .asRuntimeException());
            return;
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        PushTelemetryResponse r = PushTelemetryResponse.newBuilder().setId(request.getId()).addAllCommIds(commandIds).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void pullCommand(PullCommandRequest request, StreamObserver<PullCommandResponse> responseObserver) {
        ControlCommandDto command;
        try {
            command = this.service.pullCommand(request.getId(), request.getCommandId());
        } catch (NoSuchElementException e) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(e)
                    .withDescription(e.getMessage())
                    .asRuntimeException());
            return;
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        CommandType type = CommandType.valueOf(command.getType().toString());
        PullCommandResponse r = PullCommandResponse.newBuilder().setId(request.getId()).setCommandId(request.getCommandId()).setType(type).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}