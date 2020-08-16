package net.tomofiles.skysign.communication.api.dpo;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.ControlResponseDpo;

public class ControlResponseDpoGrpc implements ControlResponseDpo {

    private Communication communication = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    public boolean isEmpty() {
        return this.communication == null;
    }
}