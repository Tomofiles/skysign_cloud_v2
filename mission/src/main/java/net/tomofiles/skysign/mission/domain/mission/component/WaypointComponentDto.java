package net.tomofiles.skysign.mission.domain.mission.component;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;

@Data
@AllArgsConstructor
@EqualsAndHashCode
public class WaypointComponentDto {
    private int order;
    private double latitude;
    private double longitude;
    private double heightWGS84M;
    private double speedMS;
}