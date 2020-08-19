package net.tomofiles.skysign.mission.domain.mission;

import java.util.stream.Collectors;

import net.tomofiles.skysign.mission.domain.mission.component.MissionComponentDto;
import net.tomofiles.skysign.mission.domain.mission.component.WaypointComponentDto;

public class MissionFactory {
    
    public static Mission newInstance(Generator generator) {
        return new Mission(generator.newMissionId(), generator.newVersion(), generator);
    }

    public static Mission newInstance(MissionId missionId, Generator generator) {
        return new Mission(missionId, generator.newVersion(), generator);
    }

    public static Mission assembleFrom(MissionComponentDto dto, Generator generator) {
        Mission mission = new Mission(
                new MissionId(dto.getId()),
                new Version(dto.getVersion()),
                generator);

        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(dto.getTakeoffPointGroundHeightWGS84M()));
        dto.getWaypoints()
                .forEach(waypoint -> {
                        navigation.pushNextWaypoint(
                            new GeodesicCoordinates(
                                    waypoint.getLatitude(),
                                    waypoint.getLongitude()),
                            Height.distanceFrom(Height.fromM(waypoint.getHeightWGS84M()), navigation.getTakeoffPointGroundHeight()),
                            Speed.fromMS(waypoint.getSpeedMS()));
                });

        mission.setMissionName(dto.getName());
        mission.setNavigation(navigation);
        return mission;
    }

    public static MissionComponentDto takeApart(Mission mission) {
        return new MissionComponentDto(
                mission.getId().getId(),
                mission.getMissionName(),
                mission.getNavigation().getTakeoffPointGroundHeight().getHeightM(),
                mission.getVersion().getVersion(),
                mission.getNewVersion().getVersion(),
                mission.getNavigation().getWaypoints().stream()
                        .map(waypoint -> {
                            return new WaypointComponentDto(
                                    waypoint.getOrder(),
                                    waypoint.getLatitude(),
                                    waypoint.getLongitude(),
                                    Height.plus(
                                            Height.fromM(waypoint.getRelativeHeightM()),
                                            Height.fromM(mission.getNavigation().getTakeoffPointGroundHeight().getHeightM()))
                                    .getHeightM(),
                                    waypoint.getSpeedMS());
                        })
                        .collect(Collectors.toList())
                );
    }
}