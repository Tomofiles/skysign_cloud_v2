package net.tomofiles.skysign.communication.domain.vehicle;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.event.Event;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class CommunicationIdChangedEvent implements Event {
    private final CommunicationId beforeId;
    private final CommunicationId afterId;
    private final VehicleId vehicleId;
    private final Version version;

    public boolean isFirst() {
        return this.beforeId == null; // beforeIdがNullの場合、新規登録。
    }
}