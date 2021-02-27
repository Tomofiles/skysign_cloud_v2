package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.ControlResponseDpo;
import proto.skysign.ControlRequest;

@RequiredArgsConstructor
public class ControlResponseDpoGrpc implements ControlResponseDpo {

    private final ControlRequest request;
    private Communication communication = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    public boolean isEmpty() {
        return this.communication == null;
    }

    public proto.skysign.ControlResponse getGrpcResponse() {
        return proto.skysign.ControlResponse.newBuilder()
                .setId(request.getId())
                .build();
    }
}