package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

import net.tomofiles.skysign.mission.domain.mission.component.MissionComponentDto;
import net.tomofiles.skysign.mission.domain.mission.component.WaypointComponentDto;

public class ComponentMissionTests {
    
    /**
     * DTOからMissionエンティティを組み立てる。
     */
    @Test
    public void assembleIntoMissionTest() {
        MissionId id = MissionId.newId();
        String missionName = "mission name";
        Height takeoffPointGroundHeight = Height.fromM(10.0);
        Version version = Version.newVersion();

        double latitude1 = 0.0d;
        double longitude1 = 1.0d;
        double heightWGS84M1 = 2.0d;
        double speedMS1 = 4.0d;
        double latitude2 = 0.0d;
        double longitude2 = 1.0d;
        double heightWGS84M2 = 2.0d;
        double speedMS2 = 4.0d;
        double latitude3 = 0.0d;
        double longitude3 = 1.0d;
        double heightWGS84M3 = 2.0d;
        double speedMS3 = 4.0d;

        Mission mission = MissionFactory.assembleFrom(
                new MissionComponentDto(
                        id.getId(),
                        missionName,
                        takeoffPointGroundHeight.getHeightM(),
                        version.getVersion(),
                        Arrays.asList(new WaypointComponentDto[] {
                                new WaypointComponentDto(
                                        latitude1,
                                        longitude1,
                                        heightWGS84M1,
                                        speedMS1),
                                new WaypointComponentDto(
                                        latitude2,
                                        longitude2,
                                        heightWGS84M2,
                                        speedMS2),
                                new WaypointComponentDto(
                                        latitude3,
                                        longitude3,
                                        heightWGS84M3,
                                        speedMS3),
                        })
                )
        );

        Navigation expectNavigation = new Navigation();

        expectNavigation.setTakeoffPointGroundHeight(takeoffPointGroundHeight);
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude1, longitude1),
            Height.distanceFrom(Height.fromM(heightWGS84M1), takeoffPointGroundHeight),
            Speed.fromMS(speedMS1));
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude2, longitude2),
            Height.distanceFrom(Height.fromM(heightWGS84M2), takeoffPointGroundHeight),
            Speed.fromMS(speedMS2));
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude3, longitude3),
            Height.distanceFrom(Height.fromM(heightWGS84M3), takeoffPointGroundHeight),
            Speed.fromMS(speedMS3));

        assertEquals(mission.getId(), id);
        assertEquals(mission.getMissionName(), missionName);
        assertEquals(mission.getNavigation(), expectNavigation);
        assertEquals(mission.getVersion(), version);
        assertEquals(mission.getNewVersion(), version);
    }

    /**
     * MissionエンティティからDTOに分解する。
     */
    @Test
    public void takeApartMissionTest() {
        MissionId id = MissionId.newId();
        String missionName = "mission name";
        Height takeoffPointGroundHeight = Height.fromM(10.0);

        double latitude1 = 0.0d;
        double longitude1 = 1.0d;
        double heightWGS84M1 = 2.0d;
        double speedMS1 = 4.0d;
        double latitude2 = 0.0d;
        double longitude2 = 1.0d;
        double heightWGS84M2 = 2.0d;
        double speedMS2 = 4.0d;
        double latitude3 = 0.0d;
        double longitude3 = 1.0d;
        double heightWGS84M3 = 2.0d;
        double speedMS3 = 4.0d;

        Mission mission = MissionFactory.newInstance(id);
        mission.setMissionName(missionName);

        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(takeoffPointGroundHeight);
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude1, longitude1),
            Height.distanceFrom(Height.fromM(heightWGS84M1), takeoffPointGroundHeight),
            Speed.fromMS(speedMS1));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude2, longitude2),
            Height.distanceFrom(Height.fromM(heightWGS84M2), takeoffPointGroundHeight),
            Speed.fromMS(speedMS2));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude3, longitude3),
            Height.distanceFrom(Height.fromM(heightWGS84M3), takeoffPointGroundHeight),
            Speed.fromMS(speedMS3));

        mission.replaceNavigationWith(navigation);

        MissionComponentDto dto = MissionFactory.takeApart(mission);

        assertEquals(dto.getId(), id.getId());
        assertEquals(dto.getName(), missionName);
        assertEquals(dto.getTakeoffPointGroundHeightWGS84M(), takeoffPointGroundHeight.getHeightM());
        assertEquals(dto.getWaypoints().get(0).getLatitude(), latitude1);
        assertEquals(dto.getWaypoints().get(0).getLongitude(), longitude1);
        assertEquals(dto.getWaypoints().get(0).getHeightWGS84M(), heightWGS84M1);
        assertEquals(dto.getWaypoints().get(0).getSpeedMS(), speedMS1);
        assertEquals(dto.getWaypoints().get(1).getLatitude(), latitude2);
        assertEquals(dto.getWaypoints().get(1).getLongitude(), longitude2);
        assertEquals(dto.getWaypoints().get(1).getHeightWGS84M(), heightWGS84M2);
        assertEquals(dto.getWaypoints().get(1).getSpeedMS(), speedMS2);
        assertEquals(dto.getWaypoints().get(2).getLatitude(), latitude3);
        assertEquals(dto.getWaypoints().get(2).getLongitude(), longitude3);
        assertEquals(dto.getWaypoints().get(2).getHeightWGS84M(), heightWGS84M3);
        assertEquals(dto.getWaypoints().get(2).getSpeedMS(), speedMS3);
        assertEquals(dto.getVersion(), mission.getVersion().getVersion());
    }
}