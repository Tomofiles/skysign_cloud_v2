package net.tomofiles.skysign.mission.api.event.event;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class MissionCopiedWhenFlightplanCopiedEvent {
    private final String flightplanId;
    private final String originalId;
    private final String newId;
}