package net.tomofiles.skysign.communication.usecase.dpo;

import java.util.List;

import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;

public interface ListVehiclesResponsesDpo {
    public void setVehicles(List<Vehicle> vehicles);
}