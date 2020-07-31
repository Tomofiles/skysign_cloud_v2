package net.tomofiles.skysign.mission.domain.mission.component;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class WaypointComponentDto {
    private double latitude;
    private double longitude;
    private double heightWGS84M;
    private double speedMS;
}