package net.tomofiles.skysign.communication.domain.vehicle;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.Setter;
import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.event.EmptyPublisher;
import net.tomofiles.skysign.communication.event.Publisher;

@RequiredArgsConstructor(access = AccessLevel.PACKAGE)
@EqualsAndHashCode(of = {"id"})
public class Vehicle {
    @Getter
    private final VehicleId id;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private String vehicleName = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private CommunicationId commId = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version version;

    @Setter
    private Publisher publisher = new EmptyPublisher();
    
    public void nameVehicle(String name) {
        this.vehicleName = name;
    }

    public void giveCommId(CommunicationId id) {
        CommunicationId beforeId = this.commId;
        this.commId = id;
        this.publisher
                .publish(
                        new CommunicationIdChangedEvent(beforeId, id));
    }
}