package net.tomofiles.skysign.mission.api;

import net.tomofiles.skysign.mission.domain.mission.MissionId;

public class EventObjectMother {

    /**
     * テスト用MissionCopiedWhenCopiedEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalMissionCopiedWhenCopiedEvent(MissionId originalId, MissionId newId) {
        return proto.skysign.event.MissionCopiedWhenCopiedEvent.newBuilder()
            .setOriginalMissionId(originalId.getId())
            .setNewMissionId(newId.getId())
            .build()
            .toByteArray();
    }
}