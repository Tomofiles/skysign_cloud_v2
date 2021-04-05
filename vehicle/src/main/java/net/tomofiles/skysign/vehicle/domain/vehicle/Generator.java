package net.tomofiles.skysign.vehicle.domain.vehicle;

public interface Generator {
    public VehicleId newVehicleId();
    public Version newVersion();
}