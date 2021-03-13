package net.tomofiles.skysign.vehicle.api;

import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;

public class EventObjectMother {

    /**
     * テスト用VehicleCopiedWhenCopiedEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalVehicleCopiedWhenCopiedEvent(String flightplanId, VehicleId originalId, VehicleId newId) {
        return proto.skysign.event.VehicleCopiedWhenFlightplanCopiedEvent.newBuilder()
            .setFlightplanId(flightplanId)
            .setOriginalVehicleId(originalId.getId())
            .setNewVehicleId(newId.getId())
            .build()
            .toByteArray();
    }
}