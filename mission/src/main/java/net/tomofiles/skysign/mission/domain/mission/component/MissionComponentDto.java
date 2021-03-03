package net.tomofiles.skysign.mission.domain.mission.component;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;

@Data
@AllArgsConstructor
@EqualsAndHashCode
public class MissionComponentDto {
    private String id;
    private String name;
    private double takeoffPointGroundHeightWGS84M;
    private boolean isCarbonCopy;
    private String version;
    private String newVersion;
    private List<WaypointComponentDto> waypoints;
}