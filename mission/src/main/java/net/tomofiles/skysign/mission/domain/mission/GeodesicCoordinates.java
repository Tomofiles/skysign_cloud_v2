package net.tomofiles.skysign.mission.domain.mission;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class GeodesicCoordinates {
    /** 緯度（WGS84） */
    private final double latitude;
    /** 経度（WGS84） */
    private final double longitude;
}