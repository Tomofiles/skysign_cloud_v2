package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.MissionId;
import net.tomofiles.skysign.communication.service.dpo.PullUploadMissionResponseDpo;
import proto.skysign.PullUploadMissionRequest;

@RequiredArgsConstructor
public class PullUploadMissionResponseDpoGrpc implements PullUploadMissionResponseDpo {

    private final PullUploadMissionRequest request;
    private Communication communication = null;
    private MissionId missionId = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    @Override
    public void setMissionId(MissionId missionId) {
        this.missionId = missionId;
    }

    public boolean notExistCommunication() {
        return this.communication == null;
    }

    public boolean notExistCommand() {
        return this.missionId == null;
    }

    public proto.skysign.PullUploadMissionResponse getGrpcResponse() {
        return proto.skysign.PullUploadMissionResponse.newBuilder()
                .setId(request.getId())
                .setCommandId(request.getCommandId())
                .setMissionId(this.missionId.getId())
                .build();
    }
}