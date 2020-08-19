package net.tomofiles.skysign.mission.domain.mission;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@Getter
@AllArgsConstructor
@EqualsAndHashCode(of = {"order", "latitude", "longitude", "relativeHeightM", "speedMS"})
@ToString
public class Waypoint {
    /** 順序 */
    private final int order;
    /** 緯度（WGS84） */
    private final double latitude;
    /** 経度（WGS84） */
    private final double longitude;
    /** Takeoffポイントの地表高度からの高さ（m） */
    private final double relativeHeightM;
    /** ウェイポイントを通過する速度（m/s） */
    private final double speedMS;
}