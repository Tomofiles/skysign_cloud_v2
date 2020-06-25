package net.tomofiles.skysign.communication.domain.vehicle;

public interface VehicleRepository {
    void save(Vehicle vehicle);
    Vehicle getById(VehicleId id);
}