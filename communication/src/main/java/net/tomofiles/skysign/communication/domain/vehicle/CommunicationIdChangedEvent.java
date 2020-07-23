package net.tomofiles.skysign.communication.domain.vehicle;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.event.Event;

@AllArgsConstructor
@Getter
@EqualsAndHashCode(of = {"beforeId", "afterId"})
public class CommunicationIdChangedEvent implements Event {
    private final CommunicationId beforeId;
    private final CommunicationId afterId;
    private final Version version;
}