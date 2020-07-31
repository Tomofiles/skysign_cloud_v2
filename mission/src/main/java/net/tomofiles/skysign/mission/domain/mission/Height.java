package net.tomofiles.skysign.mission.domain.mission;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;

@Getter
@AllArgsConstructor(access = AccessLevel.PRIVATE)
@EqualsAndHashCode(of = "heightM")
public class Height {
    private double heightM;

    public static Height fromM(double height) {
        return new Height(height);
    }

    public static Height distanceFrom(Height airHeightOfWGS84Ellipsoid, Height groundHeightOfWGS84Ellipsoid) {
        return new Height(airHeightOfWGS84Ellipsoid.getHeightM() - groundHeightOfWGS84Ellipsoid.getHeightM());
    }

    public static Height plus(Height airHeightOfGround, Height takeoffPointGroundHeightWGS84Ellipsoid) {
        return new Height(airHeightOfGround.getHeightM() + takeoffPointGroundHeightWGS84Ellipsoid.getHeightM());
    }
}