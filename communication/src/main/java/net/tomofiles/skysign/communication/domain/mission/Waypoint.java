package net.tomofiles.skysign.communication.domain.mission;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@Getter
@AllArgsConstructor
@EqualsAndHashCode(of = {"order", "latitude", "longitude", "relativeHeightM", "speedMS"})
@ToString
public class Waypoint {
    private final int order;
    private final double latitude;
    private final double longitude;
    private final double relativeHeightM;
    private final double speedMS;
}