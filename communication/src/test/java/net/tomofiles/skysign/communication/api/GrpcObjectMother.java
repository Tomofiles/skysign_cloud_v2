package net.tomofiles.skysign.communication.api;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import proto.skysign.CreateVehicleRequest;
import proto.skysign.UpdateVehicleRequest;
import proto.skysign.Vehicle;

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

    public static CreateVehicleRequest newNormalCreateVehicleRequestGrpc() {
        return CreateVehicleRequest.newBuilder()
                .setName("vehicle name")
                .setCommId("comm id")
                .build();
    }

    public static UpdateVehicleRequest newNormalUpdateVehicleRequestGrpc(VehicleId vehicleId) {
        return UpdateVehicleRequest.newBuilder()
                .setId(vehicleId.getId())
                .setName("vehicle name")
                .setCommId("comm id")
                .build();
    }
}