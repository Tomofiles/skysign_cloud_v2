package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.PushCommandResponseDpo;
import proto.skysign.PushCommandRequest;

@RequiredArgsConstructor
public class PushCommandResponseDpoGrpc implements PushCommandResponseDpo {

    private final PushCommandRequest request;
    private Communication communication = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    public boolean isEmpty() {
        return this.communication == null;
    }

    public proto.skysign.PushCommandResponse getGrpcResponse() {
        return proto.skysign.PushCommandResponse.newBuilder()
                .setId(request.getId())
                .setType(request.getType())
                .build();
    }
}