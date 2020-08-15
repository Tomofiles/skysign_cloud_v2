package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.api.dpo.GetCommunicationRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.GetCommunicationResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PullCommandRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PullCommandResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PushTelemetryRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PushTelemetryResponseDpoGrpc;
import net.tomofiles.skysign.communication.service.CommunicateEdgeService;
import proto.skysign.common.Communication;
import proto.skysign.GetCommunicationRequest;
import proto.skysign.PullCommandRequest;
import proto.skysign.PullCommandResponse;
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
    public void getCommunication(GetCommunicationRequest request, StreamObserver<Communication> responseObserver) {
        GetCommunicationRequestDpoGrpc requestDpo = new GetCommunicationRequestDpoGrpc(request);
        GetCommunicationResponseDpoGrpc responseDpo = new GetCommunicationResponseDpoGrpc();
        
        try {
            this.service.getCommunication(requestDpo, responseDpo);
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
}