package net.tomofiles.skysign.mission.domain.mission;

public class MissionObjectMother {

    /**
     * 1件のNavigationを持つテスト用Missionエンティティを生成する。
     */
    public static Mission newSingleNavigationMission(MissionId missionId, Version version, Generator generator) {
        Mission mission = Mission.newOriginal(missionId, version, generator);
        mission.setMissionName("mission name");
        mission.setNavigation(newSingleNavigation());
        return mission;
    }

    /**
     * 1件のNavigationを持つカーボンコピーされたテスト用Missionエンティティを生成する。
     */
    public static Mission newSingleNavigationCarbonCopiedMission(MissionId missionId, Version version, Generator generator) {
        Mission mission = Mission.newCarbonCopy(missionId, version, generator);
        mission.setMissionName("mission name");
        mission.setNavigation(newSingleNavigation());
        return mission;
    }

    /**
     * 複数件のNavigationを持つテスト用Missionエンティティを生成する。
     */
    public static Mission newSeveralNavigationMission(MissionId missionId, Version version, Generator generator) {
        Mission mission = Mission.newOriginal(missionId, version, generator);
        mission.setMissionName("mission name");
        mission.setNavigation(newSeveralNavigation());
        return mission;
    }

    /**
     * 複数件のNavigationを持つカーボンコピーされたテスト用Missionエンティティを生成する。
     */
    public static Mission newSeveralNavigationCarbonCopiedMission(MissionId missionId, Version version, Generator generator) {
        Mission mission = Mission.newCarbonCopy(missionId, version, generator);
        mission.setMissionName("mission name");
        mission.setNavigation(newSeveralNavigation());
        return mission;
    }

    /**
     * 1件のNavigationオブジェクトを生成する。
     */
    public static Navigation newSingleNavigation() {
        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));
        return navigation;
    }

    /**
     * 昇順のWaypointを複数件持つNavigationオブジェクトを生成する。
     */
    public static Navigation newSeveralNavigation() {
        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(11.0, 12.0),
            Height.fromM(13.0),
            Speed.fromMS(14.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(21.0, 22.0),
            Height.fromM(23.0),
            Speed.fromMS(24.0));
        return navigation;
    }

    /**
     * 順不同のWaypointを複数件持つNavigationオブジェクトを生成する。
     */
    public static Navigation newSeveralInRondomOrderNavigation() {
        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(11.0, 12.0),
            Height.fromM(13.0),
            Speed.fromMS(14.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(21.0, 22.0),
            Height.fromM(23.0),
            Speed.fromMS(24.0));
        return navigation;
    }
}