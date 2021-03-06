package net.tomofiles.skysign.communication.domain.communication;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;
import net.tomofiles.skysign.communication.event.Event;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class TelemetryUpdatedEvent implements Event {
    private final CommunicationId communicationId;
    private final TelemetrySnapshot telemetry;
}