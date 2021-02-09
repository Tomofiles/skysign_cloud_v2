package net.tomofiles.skysign.communication.api.grpc;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.PushUploadMissionResponseDpo;

public class PushUploadMissionResponseDpoGrpc implements PushUploadMissionResponseDpo {

    private Communication communication = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    public boolean isEmpty() {
        return this.communication == null;
    }
}