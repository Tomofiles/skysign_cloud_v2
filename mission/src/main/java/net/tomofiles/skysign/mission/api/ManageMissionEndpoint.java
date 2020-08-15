package net.tomofiles.skysign.mission.api;

import java.util.NoSuchElementException;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.mission.api.dpo.CreateMissionRequestDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.CreateMissionResponseDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.DeleteMissionRequestDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.DeleteMissionResponseDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.GetMissionRequestDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.GetMissionResponseDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.ListMissionsResponsesDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.UpdateMissionRequestDpoGrpc;
import net.tomofiles.skysign.mission.api.dpo.UpdateMissionResponseDpoGrpc;
import net.tomofiles.skysign.mission.service.ManageMissionService;
import proto.skysign.DeleteMissionRequest;
import proto.skysign.common.Empty;
import proto.skysign.GetMissionRequest;
import proto.skysign.ListMissionsResponses;
import proto.skysign.common.Mission;
import proto.skysign.ManageMissionServiceGrpc.ManageMissionServiceImplBase;

@GRpcService
@Controller
@AllArgsConstructor
public class ManageMissionEndpoint extends ManageMissionServiceImplBase {

    private final ManageMissionService manageMissionService;

    @Override
    public void listMissions(Empty request, StreamObserver<ListMissionsResponses> responseObserver) {
        ListMissionsResponsesDpoGrpc responsesDpo = new ListMissionsResponsesDpoGrpc();

        try {
            this.manageMissionService.listMissions(responsesDpo);
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
    public void getMission(GetMissionRequest request, StreamObserver<Mission> responseObserver) {
        GetMissionRequestDpoGrpc requestDpo = new GetMissionRequestDpoGrpc(request);
        GetMissionResponseDpoGrpc responsesDpo = new GetMissionResponseDpoGrpc();

        try {
            this.manageMissionService.getMission(requestDpo, responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responsesDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("mission-idに合致するMissionが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responsesDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void createMission(Mission request, StreamObserver<Mission> responseObserver) {
        CreateMissionRequestDpoGrpc requestDpo = new CreateMissionRequestDpoGrpc(request);
        CreateMissionResponseDpoGrpc responsesDpo = new CreateMissionResponseDpoGrpc();

        try {
            this.manageMissionService.createMission(requestDpo, responsesDpo);
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
    public void updateMission(Mission request, StreamObserver<Mission> responseObserver) {
        UpdateMissionRequestDpoGrpc requestDpo = new UpdateMissionRequestDpoGrpc(request);
        UpdateMissionResponseDpoGrpc responsesDpo = new UpdateMissionResponseDpoGrpc();

        try {
            this.manageMissionService.updateMission(requestDpo, responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responsesDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("mission-idに合致するMissionが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responsesDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void deleteMission(DeleteMissionRequest request, StreamObserver<Empty> responseObserver) {
        DeleteMissionRequestDpoGrpc requestDpo = new DeleteMissionRequestDpoGrpc(request);
        DeleteMissionResponseDpoGrpc responsesDpo = new DeleteMissionResponseDpoGrpc();

        try {
            this.manageMissionService.deleteMission(requestDpo, responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responsesDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("mission-idに合致するMissionが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(Empty.newBuilder().build());
        responseObserver.onCompleted();
    }
}