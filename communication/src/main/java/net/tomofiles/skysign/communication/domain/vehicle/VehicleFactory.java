package net.tomofiles.skysign.communication.domain.vehicle;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public class VehicleFactory {

    public static Vehicle newInstance(VehicleId id) {
        Vehicle vehicle = new Vehicle(id);
        Version version = Version.newVersion();
        vehicle.setVersion(version);
        vehicle.setNewVersion(version);
        return vehicle;
    }

    public static Vehicle rebuild(VehicleId id, String name, String commId, String version) {
        Vehicle vehicle = new Vehicle(id);
        vehicle.setVehicleName(name);
        vehicle.setCommId(new CommunicationId(commId));
        vehicle.setVersion(new Version(version));
        vehicle.setNewVersion(new Version(version));
        return vehicle;
    }
}