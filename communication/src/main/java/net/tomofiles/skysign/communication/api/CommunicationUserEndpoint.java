package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.api.dpo.CancelRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.CancelResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.ControlRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.ControlResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.ListCommunicationsResponsesDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PullTelemetryRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PullTelemetryResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PushCommandRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.PushCommandResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.StagingRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.StagingResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.UncontrolRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.UncontrolResponseDpoGrpc;
import net.tomofiles.skysign.communication.service.CommunicationUserService;
import proto.skysign.CancelRequest;
import proto.skysign.CancelResponse;
import proto.skysign.ControlRequest;
import proto.skysign.ControlResponse;
import proto.skysign.common.Empty;
import proto.skysign.ListCommunicationsResponses;
import proto.skysign.PullTelemetryRequest;
import proto.skysign.PullTelemetryResponse;
import proto.skysign.PushCommandRequest;
import proto.skysign.PushCommandResponse;
import proto.skysign.StagingRequest;
import proto.skysign.StagingResponse;
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
        PushCommandResponseDpoGrpc responseDpo = new PushCommandResponseDpoGrpc();

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

        PushCommandResponse r = PushCommandResponse.newBuilder()
                .setId(request.getId())
                .setType(request.getType())
                .build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void staging(StagingRequest request, StreamObserver<StagingResponse> responseObserver) {
        StagingRequestDpoGrpc requestDpo = new StagingRequestDpoGrpc(request);
        StagingResponseDpoGrpc responseDpo = new StagingResponseDpoGrpc();

        try {
            this.service.staging(requestDpo, responseDpo);
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

        StagingResponse r = StagingResponse.newBuilder()
                .setMissionId(request.getMissionId())
                .setId(request.getId())
                .build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void cancel(CancelRequest request, StreamObserver<CancelResponse> responseObserver) {
        CancelRequestDpoGrpc requestDpo = new CancelRequestDpoGrpc(request);
        CancelResponseDpoGrpc responseDpo = new CancelResponseDpoGrpc();

        try {
            this.service.cancel(requestDpo, responseDpo);
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

        CancelResponse r = CancelResponse.newBuilder()
                .setId(request.getId())
                .build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void control(ControlRequest request, StreamObserver<ControlResponse> responseObserver) {
        ControlRequestDpoGrpc requestDpo = new ControlRequestDpoGrpc(request);
        ControlResponseDpoGrpc responseDpo = new ControlResponseDpoGrpc();

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

        ControlResponse r = ControlResponse.newBuilder()
                .setId(request.getId())
                .build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void uncontrol(UncontrolRequest request, StreamObserver<UncontrolResponse> responseObserver) {
        UncontrolRequestDpoGrpc requestDpo = new UncontrolRequestDpoGrpc(request);
        UncontrolResponseDpoGrpc responseDpo = new UncontrolResponseDpoGrpc();

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

        UncontrolResponse r = UncontrolResponse.newBuilder()
                .setId(request.getId())
                .build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}