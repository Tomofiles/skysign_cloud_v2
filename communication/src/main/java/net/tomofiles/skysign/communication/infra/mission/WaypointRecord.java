package net.tomofiles.skysign.communication.infra.mission;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class WaypointRecord {
    private String missionId;
    private int order;
    private double latitudeDegree;
    private double longitudeDegree;
    private double relativeHeightM;
    private double speedMS;
}