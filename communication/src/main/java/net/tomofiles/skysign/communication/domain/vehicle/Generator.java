package net.tomofiles.skysign.communication.domain.vehicle;

public interface Generator {
    public VehicleId newVehicleId();
    public Version newVersion();
}