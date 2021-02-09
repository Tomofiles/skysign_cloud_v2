package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.UncontrolResponseDpo;
import proto.skysign.UncontrolRequest;

@RequiredArgsConstructor
public class UncontrolResponseDpoGrpc implements UncontrolResponseDpo {

    private final UncontrolRequest request;
    private Communication communication = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    public boolean isEmpty() {
        return this.communication == null;
    }

    public proto.skysign.UncontrolResponse getGrpcResponse() {
        return proto.skysign.UncontrolResponse.newBuilder()
                .setId(request.getId())
                .build();
    }
}