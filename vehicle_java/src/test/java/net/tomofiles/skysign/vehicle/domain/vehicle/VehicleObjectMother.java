package net.tomofiles.skysign.vehicle.domain.vehicle;

public class VehicleObjectMother {
    
    /**
     * テスト用Vehicleエンティティを生成する。
     */
    public static Vehicle newNormalVehicle(VehicleId vehicleId, Version version, Generator generator) {
        Vehicle vehicle = Vehicle.newOriginal(vehicleId, version, generator);
        vehicle.setVehicleName("vehicle name");
        vehicle.setCommunicationId(new CommunicationId("comm id"));
        return vehicle;
    }

    /**
     * テスト用のカーボンコピーされたVehicleエンティティを生成する。
     */
    public static Vehicle newCarbonCopiedVehicle(VehicleId vehicleId, Version version, Generator generator) {
        Vehicle vehicle = Vehicle.newCarbonCopy(vehicleId, version, generator);
        vehicle.setVehicleName("vehicle name");
        vehicle.setCommunicationId(new CommunicationId("comm id"));
        return vehicle;
    }
}