package net.tomofiles.skysign.communication.api.event.event;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class CommunicationIdGaveEvent {
    private final String communicationId;
    private final String vehicleId;
    private final String version;
}