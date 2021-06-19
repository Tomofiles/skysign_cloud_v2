package net.tomofiles.skysign.mission.domain.mission;

import java.util.ArrayList;
import java.util.List;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@EqualsAndHashCode(of = {"takeoffPointGroundHeight", "waypoints"})
@ToString
public class Navigation {

    /** 現在の順序 */
    private int currentOrder;

    /** Takeoffポイントの地表高度（WGS84楕円体高） */
    @Getter
    @Setter
    private Height takeoffPointGroundHeight;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private List<Waypoint> waypoints;

    public Navigation() {
        this.currentOrder = 1;
        this.waypoints = new ArrayList<>();
    }
    
    public void pushNextWaypoint(GeodesicCoordinates coordinates, Height relativeHeight, Speed speed) {
        this.waypoints.add(
            new Waypoint(
                this.currentOrder++,
                coordinates.getLatitude(),
                coordinates.getLongitude(),
                relativeHeight.getHeightM(),
                speed.getSpeedMS()
            )
        );
    }
}