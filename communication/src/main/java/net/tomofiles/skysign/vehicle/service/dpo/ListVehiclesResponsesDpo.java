package net.tomofiles.skysign.vehicle.service.dpo;

import java.util.List;

import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;

public interface ListVehiclesResponsesDpo {
    public void setVehicles(List<Vehicle> vehicles);
}