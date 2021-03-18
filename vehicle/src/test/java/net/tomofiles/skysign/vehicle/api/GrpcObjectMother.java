package net.tomofiles.skysign.vehicle.api;

import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationId;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import proto.skysign.common.Vehicle;

public class GrpcObjectMother {

    /**
     * テスト用Vehicleオブジェクトを生成する。
     */
    public static Vehicle newNormalVehicleGrpc(VehicleId vehicleId) {
        return Vehicle.newBuilder()
                .setId(vehicleId.getId())
                .setName("vehicle name")
                .setCommId(new CommunicationId("comm id").getId())
                .build();
    }

    /**
     * Vehicle idが無いテスト用Vehicleオブジェクトを生成する。
     */
    public static Vehicle newNoIdVehicleGrpc() {
        return Vehicle.newBuilder()
                .setName("vehicle name")
                .setCommId(new CommunicationId("comm id").getId())
                .build();
    }
}