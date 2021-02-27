package net.tomofiles.skysign.vehicle.domain.vehicle;

public class VehicleObjectMother {
    
    /**
     * テスト用Vehicleエンティティを生成する。
     */
    public static Vehicle newNormalVehicle(VehicleId vehicleId, Version version, Generator generator) {
        Vehicle vehicle = new Vehicle(vehicleId, version, generator);
        vehicle.setVehicleName("vehicle name");
        vehicle.setCommId(new CommunicationId("comm id"));
        return vehicle;
    }
}