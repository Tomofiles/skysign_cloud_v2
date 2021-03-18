package net.tomofiles.skysign.vehicle.domain.vehicle;

public class VehicleFactory {

    public static Vehicle newInstance(Generator generator) {
        return Vehicle.newOriginal(generator.newVehicleId(), generator.newVersion(), generator);
    }

    public static Vehicle copy(Vehicle original, VehicleId newId, Generator generator) {
        Vehicle vehicle = Vehicle.newCarbonCopy(newId, original.getVersion(), generator);
        vehicle.setVehicleName(original.getVehicleName());
        vehicle.setCommId(original.getCommId());
        return vehicle;
    }

    public static Vehicle rebuild(VehicleId id, String name, String commId, boolean isCarbonCopy, String version, Generator generator) {
        Vehicle vehicle = isCarbonCopy
                ? Vehicle.newCarbonCopy(id, new Version(version), generator)
                : Vehicle.newOriginal(id, new Version(version), generator);
        vehicle.setVehicleName(name);
        vehicle.setCommId(new CommunicationId(commId));
        return vehicle;
    }
}