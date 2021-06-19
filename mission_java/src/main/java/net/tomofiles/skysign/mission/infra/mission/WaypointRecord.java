package net.tomofiles.skysign.mission.infra.mission;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class WaypointRecord {
    private String missionId;
    private int order;
    private double latitude;
    private double longitude;
    private double heightWGS84EllipsoidM;
    private double speedMS;
}