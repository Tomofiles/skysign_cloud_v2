package net.tomofiles.skysign.communication.domain.vehicle;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.event.EmptyPublisher;
import net.tomofiles.skysign.communication.event.Publisher;

@EqualsAndHashCode(of = {"id"})
@ToString
public class Vehicle {
    @Getter
    private final VehicleId id;

    private final Generator generator;
    
    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private String vehicleName = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private CommunicationId commId = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version version;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version newVersion;

    @Setter
    private Publisher publisher = new EmptyPublisher();
    
    Vehicle(VehicleId id, Version version, Generator generator) {
        this.id = id;
        this.version = version;
        this.newVersion = version;

        this.generator = generator;
    }

    public void nameVehicle(String name) {
        this.vehicleName = name;
        this.newVersion = this.generator.newVersion();
    }

    public void giveCommId(CommunicationId id) {
        CommunicationId beforeId = this.commId;
        this.commId = id;
        this.newVersion = this.generator.newVersion();
        this.publisher
                .publish(
                        new CommunicationIdChangedEvent(beforeId, id, this.id, this.newVersion));
    }
}