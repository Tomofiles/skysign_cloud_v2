package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.PushUploadMissionResponseDpo;
import proto.skysign.PushUploadMissionRequest;

@RequiredArgsConstructor
public class PushUploadMissionResponseDpoGrpc implements PushUploadMissionResponseDpo {

    private final PushUploadMissionRequest request;
    private Communication communication = null;
    private CommandId commandId = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    @Override
    public void setCommandId(CommandId commandId) {
        this.commandId = commandId;
    }

    public boolean isEmpty() {
        return this.communication == null;
    }

    public proto.skysign.PushUploadMissionResponse getGrpcResponse() {
        return proto.skysign.PushUploadMissionResponse.newBuilder()
                .setId(request.getId())
                .setMissionId(request.getMissionId())
                .setCommandId(this.commandId.getId())
                .build();
    }
}