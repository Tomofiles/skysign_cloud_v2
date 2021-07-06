package net.tomofiles.skysign.communication.domain.mission;

import java.util.ArrayList;
import java.util.List;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@EqualsAndHashCode(of = {"id"})
@ToString
public class Mission {
    @Getter
    private final MissionId id;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private List<Waypoint> waypoints;

    public Mission(MissionId id) {
        this.id = id;
        this.waypoints = new ArrayList<>();
    }

    public int pushWaypoint(
        double latitudeDegree, 
        double longitudeDegree, 
        double relativeHeightM, 
        double speedMS) {
            int order = this.waypoints.size() + 1;
            this.waypoints.add(
                new Waypoint(
                    order, 
                latitudeDegree, 
                longitudeDegree,
                 relativeHeightM, 
                 speedMS)
            );
            return order;
    }
}