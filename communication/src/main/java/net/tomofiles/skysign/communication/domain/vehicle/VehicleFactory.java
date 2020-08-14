package net.tomofiles.skysign.communication.domain.vehicle;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public class VehicleFactory {

    public static Vehicle newInstance(Generator generator) {
        return new Vehicle(generator.newVehicleId(), generator.newVersion(), generator);
    }

    public static Vehicle rebuild(VehicleId id, String name, String commId, String version, Generator generator) {
        Vehicle vehicle = new Vehicle(id, new Version(version), generator);
        vehicle.setVehicleName(name);
        vehicle.setCommId(new CommunicationId(commId));
        return vehicle;
    }
}