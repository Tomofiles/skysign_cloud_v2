package net.tomofiles.skysign.communication.api;

import java.util.NoSuchElementException;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.usecase.ControlVehicleService;
import net.tomofiles.skysign.communication.usecase.dto.ControlCommandType;
import net.tomofiles.skysign.communication.usecase.dto.TelemetryDto;
import proto.skysign.PullTelemetryRequest;
import proto.skysign.PullTelemetryResponse;
import proto.skysign.PushCommandRequest;
import proto.skysign.PushCommandResponse;
import proto.skysign.CommunicationUserServiceGrpc.CommunicationUserServiceImplBase;

@GRpcService
@Controller
public class ControlVehicleEndpoint extends CommunicationUserServiceImplBase {
    
    @Autowired
    private ControlVehicleService service;

    @Override
    public void pullTelemetry(PullTelemetryRequest request, StreamObserver<PullTelemetryResponse> responseObserver) {
        TelemetryDto telemetry;
        
        try {
            telemetry = this.service.pullTelemetry(request.getId());
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

        PullTelemetryResponse r = PullTelemetryResponse.newBuilder()
                .setName(telemetry.getName())
                .setLatitude(telemetry.getLatitude())
                .setLongitude(telemetry.getLongitude())
                .setAltitude(telemetry.getAltitude())
                .setRelativeAltitude(telemetry.getRelativeAltitude())
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
        try {
            this.service.pushCommand(request.getId(), ControlCommandType.valueOf(request.getType().name()));
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

        PushCommandResponse r = PushCommandResponse.newBuilder().setId(request.getId()).setType(request.getType()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}