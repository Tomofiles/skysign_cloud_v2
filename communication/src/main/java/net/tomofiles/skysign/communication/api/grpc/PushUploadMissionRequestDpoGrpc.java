package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.MissionId;
import net.tomofiles.skysign.communication.service.dpo.PushUploadMissionRequestDpo;
import proto.skysign.PushUploadMissionRequest;

@RequiredArgsConstructor
public class PushUploadMissionRequestDpoGrpc implements PushUploadMissionRequestDpo {

    private final PushUploadMissionRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }

    @Override
    public MissionId getMissionId() {
        return new MissionId(this.request.getMissionId());
    }
}