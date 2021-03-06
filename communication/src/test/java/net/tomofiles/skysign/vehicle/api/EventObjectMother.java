package net.tomofiles.skysign.vehicle.api;

import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;

public class EventObjectMother {

    /**
     * テスト用VehicleCopiedWhenCopiedEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalVehicleCopiedWhenCopiedEvent(VehicleId originalId, VehicleId newId) {
        return proto.skysign.event.VehicleCopiedWhenCopiedEvent.newBuilder()
            .setOriginalVehicleId(originalId.getId())
            .setNewVehicleId(newId.getId())
            .build()
            .toByteArray();
    }
}