package net.tomofiles.skysign.communication.api.event.event;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class CopiedMissionCreatedEvent {
    private final String missionId;
    private final List<Waypoint> waypoints;
}