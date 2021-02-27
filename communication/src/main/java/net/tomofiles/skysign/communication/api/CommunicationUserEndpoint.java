package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.api.grpc.ControlRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.ControlResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.ListCommunicationsResponsesDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PullTelemetryRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PullTelemetryResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PushCommandRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PushCommandResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PushUploadMissionRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.PushUploadMissionResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.UncontrolRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.UncontrolResponseDpoGrpc;
import net.tomofiles.skysign.communication.service.CommunicationUserService;
import proto.skysign.ControlRequest;
import proto.skysign.ControlResponse;
import proto.skysign.common.Empty;
import proto.skysign.ListCommunicationsResponses;
import proto.skysign.PullTelemetryRequest;
import proto.skysign.PullTelemetryResponse;
import proto.skysign.PushCommandRequest;
import proto.skysign.PushCommandResponse;
import proto.skysign.PushUploadMissionRequest;
import proto.skysign.PushUploadMissionResponse;
import proto.skysign.UncontrolRequest;
import proto.skysign.UncontrolResponse;
import proto.skysign.CommunicationUserServiceGrpc.CommunicationUserServiceImplBase;

@GRpcService
@Controller
@AllArgsConstructor
public class CommunicationUserEndpoint extends CommunicationUserServiceImplBase {
    
    private final CommunicationUserService service;

    @Override
    public void listCommunications(Empty request, StreamObserver<ListCommunicationsResponses> responseObserver) {
        ListCommunicationsResponsesDpoGrpc responsesDpo = new ListCommunicationsResponsesDpoGrpc();

        try {
            this.service.listCommunications(responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responsesDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void pullTelemetry(PullTelemetryRequest request, StreamObserver<PullTelemetryResponse> responseObserver) {
        PullTelemetryRequestDpoGrpc requestDpo = new PullTelemetryRequestDpoGrpc(request);
        PullTelemetryResponseDpoGrpc responseDpo  = new PullTelemetryResponseDpoGrpc();

        try {
            this.service.pullTelemetry(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するcommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void pushCommand(PushCommandRequest request, StreamObserver<PushCommandResponse> responseObserver) {
        PushCommandRequestDpoGrpc requestDpo = new PushCommandRequestDpoGrpc(request);
        PushCommandResponseDpoGrpc responseDpo = new PushCommandResponseDpoGrpc(request);

        try {
            this.service.pushCommand(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するcommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void pushUploadMission(PushUploadMissionRequest request, StreamObserver<PushUploadMissionResponse> responseObserver) {
        PushUploadMissionRequestDpoGrpc requestDpo = new PushUploadMissionRequestDpoGrpc(request);
        PushUploadMissionResponseDpoGrpc responseDpo = new PushUploadMissionResponseDpoGrpc(request);

        try {
            this.service.pushUploadMission(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するcommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void control(ControlRequest request, StreamObserver<ControlResponse> responseObserver) {
        ControlRequestDpoGrpc requestDpo = new ControlRequestDpoGrpc(request);
        ControlResponseDpoGrpc responseDpo = new ControlResponseDpoGrpc(request);

        try {
            this.service.control(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するcommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void uncontrol(UncontrolRequest request, StreamObserver<UncontrolResponse> responseObserver) {
        UncontrolRequestDpoGrpc requestDpo = new UncontrolRequestDpoGrpc(request);
        UncontrolResponseDpoGrpc responseDpo = new UncontrolResponseDpoGrpc(request);

        try {
            this.service.uncontrol(requestDpo, responseDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responseDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withDescription("communication-idに合致するcommunicationが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }
}