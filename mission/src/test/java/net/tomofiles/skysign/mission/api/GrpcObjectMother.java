package net.tomofiles.skysign.mission.api;

import net.tomofiles.skysign.mission.domain.mission.Generator;
import proto.skysign.Mission;
import proto.skysign.MissionItem;

public class GrpcObjectMother {
    
    /**
     * 1件のItemを持つテスト用Missionオブジェクトを生成する。
     */
    public static Mission newSingleItemMission(Generator generator) {
        Mission mission = Mission.newBuilder()
                .setId(generator.newMissionId().getId())
                .setName("mission name")
                .setTakeoffPointGroundHeight(0.0)
                .addItems(newSingleItem())
                .build();
        return mission;
    }

    /**
     * 1件のItemを持つID未設定のテスト用Missionオブジェクトを生成する。
     */
    public static Mission newSingleItemMissionNoID() {
        Mission mission = Mission.newBuilder()
                .setName("mission name")
                .setTakeoffPointGroundHeight(0.0)
                .addItems(newSingleItem())
                .build();
        return mission;
    }

    /**
     * 1件のItemオブジェクトを生成する。
     */
    public static MissionItem newSingleItem() {
        MissionItem item = MissionItem.newBuilder()
                .setLatitude(1.0)
                .setLongitude(2.0)
                .setRelativeHeight(3.0)
                .setSpeed(4.0)
                .build();
        return item;
    }

}