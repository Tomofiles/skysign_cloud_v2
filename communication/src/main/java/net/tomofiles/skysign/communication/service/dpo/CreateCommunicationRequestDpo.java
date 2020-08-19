package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;

public interface CreateCommunicationRequestDpo {
    public CommunicationId getCommId();
    public VehicleId getVehicleId();
}