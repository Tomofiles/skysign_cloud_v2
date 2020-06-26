package net.tomofiles.skysign.communication.domain.vehicle;

import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public class VehicleFactory {

    public static Vehicle newInstance(VehicleId id) {
        Vehicle vehicle = new Vehicle(id);
        vehicle.setVersion(new Version(1));
        return vehicle;
    }

    public static Vehicle rebuild(VehicleId id, String name, String commId, int version) {
        Vehicle vehicle = new Vehicle(id);
        vehicle.setVehicleName(name);
        vehicle.setCommId(new CommunicationId(commId));
        vehicle.setVersion(new Version(version));
        return vehicle;
    }
}