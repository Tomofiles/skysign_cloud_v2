package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.PullUploadMissionRequestDpo;
import proto.skysign.PullUploadMissionRequest;

@RequiredArgsConstructor
public class PullUploadMissionRequestDpoGrpc implements PullUploadMissionRequestDpo {

    private final PullUploadMissionRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }

    @Override
    public CommandId getCommandId() {
        return new CommandId(this.request.getCommandId());
    }
}