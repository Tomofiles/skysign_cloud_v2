package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public interface CreateVehicleRequestDpo {
    public String getVehicleName();
    public CommunicationId getCommId();
}