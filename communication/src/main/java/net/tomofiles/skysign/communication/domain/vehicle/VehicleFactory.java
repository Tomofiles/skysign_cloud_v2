package net.tomofiles.skysign.communication.domain.vehicle;

import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public class VehicleFactory {

    public static Vehicle newInstance(VehicleId id) {
        return new Vehicle(id);
    }

    public static Vehicle rebuild(VehicleId id, String name, String commId, int version) {
        Vehicle vehicle = new Vehicle(id);
        vehicle.setVehicleName(name);
        vehicle.setCommId(new CommunicationId(commId));
        vehicle.setVersion(new Version(version));
        return vehicle;
    }
}