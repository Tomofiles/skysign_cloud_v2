package net.tomofiles.skysign.vehicle.service.dpo;

import net.tomofiles.skysign.vehicle.domain.vehicle.FlightplanId;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;

public interface CarbonCopyVehicleRequestDpo {
    public VehicleId getOriginalId();
    public VehicleId getNewId();
    public FlightplanId getFlightplanId();
}