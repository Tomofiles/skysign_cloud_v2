package net.tomofiles.skysign.vehicle.service.dpo;

import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationId;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;

public interface UpdateVehicleRequestDpo {
    public VehicleId getId();
    public String getVehicleName();
    public CommunicationId getCommunicationId();
}