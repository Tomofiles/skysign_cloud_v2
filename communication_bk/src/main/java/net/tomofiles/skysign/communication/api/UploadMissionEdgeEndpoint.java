package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.api.grpc.GetUploadMissionRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.grpc.GetUploadMissionResponseDpoGrpc;
import net.tomofiles.skysign.communication.service.UploadMissionEdgeService;
import proto.skysign.GetUploadMissionRequest;
import proto.skysign.UploadMission;
import proto.skysign.UploadMissionEdgeServiceGrpc.UploadMissionEdgeServiceImplBase;

@GRpcService
@Controller
@AllArgsConstructor
public class UploadMissionEdgeEndpoint extends UploadMissionEdgeServiceImplBase {

    private final UploadMissionEdgeService service;

    @Override
    public void getUploadMission(GetUploadMissionRequest request, StreamObserver<UploadMission> responseObserver) {
        GetUploadMissionRequestDpoGrpc requestDpo = new GetUploadMissionRequestDpoGrpc(request);
        GetUploadMissionResponseDpoGrpc responseDpo  = new GetUploadMissionResponseDpoGrpc();

        try {
            this.service.getUploadMission(requestDpo, responseDpo);
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
                    .withDescription("mission-idに合致するupload missionが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responseDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }
}