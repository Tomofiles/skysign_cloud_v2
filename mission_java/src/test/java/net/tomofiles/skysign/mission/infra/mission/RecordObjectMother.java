package net.tomofiles.skysign.mission.infra.mission;

import java.util.Arrays;
import java.util.List;

import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.Version;

public class RecordObjectMother {

    /**
     * 通常のMissionレコードを生成する。
     */
    public static MissionRecord newNormalMissionRecord(MissionId missionId, Version version) {
        return new MissionRecord(
                missionId.getId(),
                "mission name",
                0.0,
                false,
                version.getVersion(),
                version.getVersion());
    }

    /**
     * カーボンコピーされたMissionレコードを生成する。
     */
    public static MissionRecord newNormalCarbonCopiedMissionRecord(MissionId missionId, Version version) {
        return new MissionRecord(
                missionId.getId(),
                "mission name",
                0.0,
                true,
                version.getVersion(),
                version.getVersion());
    }

    /**
     * 1件のWaypointレコードを生成する。
     */
    public static WaypointRecord newSingleWaypointRecord(MissionId missionId) {
        return new WaypointRecord(
                missionId.getId(),
                1,
                1.0,
                2.0,
                3.0,
                4.0);
    }

    /**
     * 複数件の昇順のWaypointレコードを生成する。
     */
    public static List<WaypointRecord> newSeveralWaypointRecords(MissionId missionId) {
        WaypointRecord waypointRecord1 = new WaypointRecord(
                missionId.getId(),
                1,
                1.0,
                2.0,
                3.0,
                4.0);
        WaypointRecord waypointRecord2 = new WaypointRecord(
                missionId.getId(),
                2,
                11.0,
                12.0,
                13.0,
                14.0);
        WaypointRecord waypointRecord3 = new WaypointRecord(
                missionId.getId(),
                3,
                21.0,
                22.0,
                23.0,
                24.0);
        return Arrays.asList(new WaypointRecord[] {
            waypointRecord1,
            waypointRecord2,
            waypointRecord3
        });
    }

    /**
     * 複数件の順不同のWaypointレコードを生成する。
     */
    public static List<WaypointRecord> newSeveralInRondomOrderWaypointRecords(MissionId missionId) {
        WaypointRecord waypointRecord1 = new WaypointRecord(
                missionId.getId(),
                1,
                1.0,
                2.0,
                3.0,
                4.0);
        WaypointRecord waypointRecord2 = new WaypointRecord(
                missionId.getId(),
                2,
                11.0,
                12.0,
                13.0,
                14.0);
        WaypointRecord waypointRecord3 = new WaypointRecord(
                missionId.getId(),
                3,
                21.0,
                22.0,
                23.0,
                24.0);
        return Arrays.asList(new WaypointRecord[] {
            waypointRecord3, // 順序がバラバラ
            waypointRecord1, // 順序がバラバラ
            waypointRecord2  // 順序がバラバラ
        });
    }

}