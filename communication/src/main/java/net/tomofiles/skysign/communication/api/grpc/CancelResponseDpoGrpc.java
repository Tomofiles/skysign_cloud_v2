package net.tomofiles.skysign.communication.api.grpc;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.CancelResponseDpo;

public class CancelResponseDpoGrpc implements CancelResponseDpo {

    private Communication communication = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    public boolean isEmpty() {
        return this.communication == null;
    }
}