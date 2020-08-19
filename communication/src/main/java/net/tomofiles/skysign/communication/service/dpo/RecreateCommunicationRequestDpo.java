package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;

public interface RecreateCommunicationRequestDpo {
    public CommunicationId getBeforeCommId();
    public CommunicationId getAfterCommId();
    public VehicleId getVehicleId();
}