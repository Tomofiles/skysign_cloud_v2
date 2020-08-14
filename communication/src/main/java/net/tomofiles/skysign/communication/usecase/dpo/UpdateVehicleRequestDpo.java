package net.tomofiles.skysign.communication.usecase.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;

public interface UpdateVehicleRequestDpo {
    public VehicleId getVehicleId();
    public String getVehicleName();
    public CommunicationId getCommId();
}