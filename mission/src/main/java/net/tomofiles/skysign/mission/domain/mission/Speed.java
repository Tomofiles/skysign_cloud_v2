package net.tomofiles.skysign.mission.domain.mission;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.ToString;

@Getter
@AllArgsConstructor(access = AccessLevel.PRIVATE)
@ToString
public class Speed {
    private double speedMS;

    public static Speed fromMS(double speed) {
        return new Speed(speed);
    }
}