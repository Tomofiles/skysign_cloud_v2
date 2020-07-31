package net.tomofiles.skysign.mission.domain.mission;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;

@Getter
@AllArgsConstructor
@EqualsAndHashCode(of = {"latitude", "longitude", "relativeHeightM", "speedMS"})
public class Waypoint {
    /** 緯度（WGS84） */
    private final double latitude;
    /** 経度（WGS84） */
    private final double longitude;
    /** Takeoffポイントの地表高度からの高さ（m） */
    private final double relativeHeightM;
    /** ウェイポイントを通過する速度（m/s） */
    private final double speedMS;
}