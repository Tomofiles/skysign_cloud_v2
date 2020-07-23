package net.tomofiles.skysign.communication.domain.vehicle;

import java.util.List;

public interface VehicleRepository {
    void save(Vehicle vehicle);
    Vehicle getById(VehicleId id);
    List<Vehicle> getAll();
    void remove(VehicleId id, Version version);
}