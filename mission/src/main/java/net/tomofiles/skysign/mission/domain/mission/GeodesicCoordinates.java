package net.tomofiles.skysign.mission.domain.mission;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.ToString;

@Getter
@AllArgsConstructor
@ToString
public class GeodesicCoordinates {
    /** 緯度（WGS84） */
    private final double latitude;
    /** 経度（WGS84） */
    private final double longitude;
}