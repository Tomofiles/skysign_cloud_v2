package net.tomofiles.skysign.mission.infra.mission;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class MissionRecord {
    private String id;
    private String name;
    private double takeoffPointGroundHeightWGS84EllipsoidM;
    private String version;
    private String newVersion;
}