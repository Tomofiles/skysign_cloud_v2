package net.tomofiles.skysign.mission.api;

import net.tomofiles.skysign.mission.domain.mission.MissionId;

public class EventObjectMother {

    /**
     * テスト用MissionCopiedWhenCopiedEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalMissionCopiedWhenCopiedEvent(String flightplanId, MissionId originalId, MissionId newId) {
        return proto.skysign.event.MissionCopiedWhenFlightplanCopiedEvent.newBuilder()
            .setFlightplanId(flightplanId)
            .setOriginalMissionId(originalId.getId())
            .setNewMissionId(newId.getId())
            .build()
            .toByteArray();
    }
}