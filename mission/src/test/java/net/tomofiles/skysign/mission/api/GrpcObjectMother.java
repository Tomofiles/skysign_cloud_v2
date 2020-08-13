package net.tomofiles.skysign.mission.api;

import net.tomofiles.skysign.mission.domain.mission.MissionId;
import proto.skysign.Mission;
import proto.skysign.MissionItem;

public class GrpcObjectMother {
    
    /**
     * 1件のItemを持つテスト用Missionオブジェクトを生成する。
     */
    public static Mission newSingleItemMissionGrpc(MissionId missionId) {
        Mission mission = Mission.newBuilder()
                .setId(missionId.getId())
                .setName("mission name")
                .setTakeoffPointGroundHeight(0.0)
                .addItems(newSingleItemGrpc())
                .build();
        return mission;
    }

    /**
     * 1件のItemを持つID未設定のテスト用Missionオブジェクトを生成する。
     */
    public static Mission newSingleItemMissionNoIDGrpc() {
        Mission mission = Mission.newBuilder()
                .setName("mission name")
                .setTakeoffPointGroundHeight(0.0)
                .addItems(newSingleItemGrpc())
                .build();
        return mission;
    }

    /**
     * 1件のItemオブジェクトを生成する。
     */
    public static MissionItem newSingleItemGrpc() {
        MissionItem item = MissionItem.newBuilder()
                .setLatitude(1.0)
                .setLongitude(2.0)
                .setRelativeHeight(3.0)
                .setSpeed(4.0)
                .build();
        return item;
    }

}