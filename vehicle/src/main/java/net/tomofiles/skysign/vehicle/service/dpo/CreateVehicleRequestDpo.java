package net.tomofiles.skysign.vehicle.service.dpo;

import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationId;

public interface CreateVehicleRequestDpo {
    public String getVehicleName();
    public CommunicationId getCommId();
}