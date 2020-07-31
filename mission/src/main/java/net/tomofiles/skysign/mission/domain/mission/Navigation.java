package net.tomofiles.skysign.mission.domain.mission;

import java.util.ArrayList;
import java.util.List;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;

@EqualsAndHashCode(of = {"takeoffPointGroundHeight", "waypoints"})
public class Navigation {

    /** Takeoffポイントの地表高度（WGS84楕円体高） */
    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Height takeoffPointGroundHeight;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private List<Waypoint> waypoints;

    public Navigation() {
        this.waypoints = new ArrayList<>();
    }
    
    public void pushNextWaypoint(GeodesicCoordinates coordinates, Height relativeHeight, Speed speed) {
        this.waypoints.add(
            new Waypoint(
                coordinates.getLatitude(),
                coordinates.getLongitude(),
                relativeHeight.getHeightM(),
                speed.getSpeedMS()
            )
        );
    }
}