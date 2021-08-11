package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.api.grpc.PullCommandRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PullCommandResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PullUploadMissionRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PullUploadMissionResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PushTelemetryRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PushTelemetryResponseDpoGrpc;
import net.tomofiles.skysign.communication.service.CommunicateEdgeService;
import proto.skysign.PullCommandRequest;
import proto.skysign.PullCommandResponse;
import proto.skysign.PullUploadMissionRequest;
import proto.skysign.PullUploadMissionResponse;
import proto.skysign.PushTelemetryRequest;
import proto.skysign.PushTelemetryResponse;
import proto.skysign.CommunicationEdgeServiceGrpc.CommunicationEdgeServiceImplBase;

@GRpcService
@Controller
@AllArgsConstructor
public class CommunicateEdgeEndpoint extends CommunicationEdgeServiceImplBase {

    private final CommunicateEdgeService service;

    @Override
    public void pushTelemetry(PushTelemetryRequest request, StreamObserver<PushTelemetryResponse> responseObserver) {
        PushTelemetryRequestDpoGrpc requestDpo = new PushTelemetryRequestDpoGrpc(request);
        PushTelemetryResponseDpoGrpc responseDpo = new PushTelemetryResponseDpoGrpc(request);

        try {
            this.service.pushTelemetry(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.notExistCommunication()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するCommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse()); 
        responseObserver.onCompleted();
    }

    @Override
    public void pullCommand(PullCommandRequest request, StreamObserver<PullCommandResponse> responseObserver) {
        PullCommandRequestDpoGrpc requestDpo = new PullCommandRequestDpoGrpc(request);
        PullCommandResponseDpoGrpc responseDpo = new PullCommandResponseDpoGrpc(request);
        
        try {
            this.service.pullCommand(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.notExistCommunication()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するCommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        if (responseDpo.notExistCommand()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("command-idに合致するCommandが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse()); 
        responseObserver.onCompleted();
    }

    @Override
    public void pullUploadMission(PullUploadMissionRequest request, StreamObserver<PullUploadMissionResponse> responseObserver) {
        PullUploadMissionRequestDpoGrpc requestDpo = new PullUploadMissionRequestDpoGrpc(request);
        PullUploadMissionResponseDpoGrpc responseDpo = new PullUploadMissionResponseDpoGrpc(request);
        
        try {
            this.service.pullUploadMission(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.notExistCommunication()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するCommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        if (responseDpo.notExistCommand()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("command-idに合致するUploadMissionが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse()); 
        responseObserver.onCompleted();
    }
}