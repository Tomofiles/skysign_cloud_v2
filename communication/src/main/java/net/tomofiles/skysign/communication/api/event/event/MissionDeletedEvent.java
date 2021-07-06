package net.tomofiles.skysign.communication.api.event.event;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class MissionDeletedEvent {
    private final String missionId;
}