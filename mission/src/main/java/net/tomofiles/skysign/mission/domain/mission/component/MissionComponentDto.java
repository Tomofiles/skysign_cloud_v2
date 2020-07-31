package net.tomofiles.skysign.mission.domain.mission.component;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class MissionComponentDto {
    private String id;
    private String name;
    private double takeoffPointGroundHeightWGS84M;
    private String version;
    private List<WaypointComponentDto> waypoints;
}