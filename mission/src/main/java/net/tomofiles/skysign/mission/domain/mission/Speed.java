package net.tomofiles.skysign.mission.domain.mission;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor(access = AccessLevel.PRIVATE)
public class Speed {
    private double speedMS;

    public static Speed fromMS(double speed) {
        return new Speed(speed);
    }
}