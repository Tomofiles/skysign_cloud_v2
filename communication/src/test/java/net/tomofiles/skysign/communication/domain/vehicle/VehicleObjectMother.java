package net.tomofiles.skysign.communication.domain.vehicle;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

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