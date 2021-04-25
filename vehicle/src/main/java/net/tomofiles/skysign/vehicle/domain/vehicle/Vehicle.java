package net.tomofiles.skysign.vehicle.domain.vehicle;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;
import net.tomofiles.skysign.vehicle.event.EmptyPublisher;
import net.tomofiles.skysign.vehicle.event.Publisher;

@EqualsAndHashCode(of = {"id"})
@ToString
public class Vehicle {
    @Getter
    private final VehicleId id;

    private final Generator generator;
    
    @Getter
    private final boolean isCarbonCopy;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private String vehicleName = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private CommunicationId communicationId = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version version;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version newVersion;

    @Setter
    private Publisher publisher = new EmptyPublisher();
    
    Vehicle(VehicleId id, boolean isCarbonCopy, Version version, Generator generator) {
        this.id = id;
        this.isCarbonCopy = isCarbonCopy;
        this.version = version;
        this.newVersion = version;

        this.generator = generator;
    }

    static Vehicle newOriginal(VehicleId id, Version version, Generator generator) {
        return new Vehicle(id, false, version, generator);
    } 

    static Vehicle newCarbonCopy(VehicleId id, Version version, Generator generator) {
        return new Vehicle(id, true, version, generator);
    } 

    public void nameVehicle(String name) {
        if (this.isCarbonCopy) {
            throw new CannotChangeVehicleException("cannot change carbon copied vehicle");
        }

        this.vehicleName = name;
        this.newVersion = this.generator.newVersion();
    }

    public void giveCommunication(CommunicationId id) {
        if (this.isCarbonCopy) {
            throw new CannotChangeVehicleException("cannot change carbon copied vehicle");
        }

        if (this.communicationId == null) {
            this.communicationId = id;
            this.newVersion = this.generator.newVersion();
            this.publisher
                    .publish(
                            new CommunicationIdGaveEvent(this.communicationId));
        } else {
            CommunicationId beforeId = this.communicationId;
            this.communicationId = id;
            this.newVersion = this.generator.newVersion();
            this.publisher
                    .publish(
                            new CommunicationIdRemovedEvent(beforeId));
            this.publisher
                    .publish(
                            new CommunicationIdGaveEvent(this.communicationId));
        }
    }

    public void removeCommunication() {
        if (this.isCarbonCopy) {
            throw new CannotChangeVehicleException("cannot change carbon copied vehicle");
        }

        CommunicationId removedId = this.communicationId;
        this.communicationId = null;
        this.newVersion = this.generator.newVersion();
        this.publisher
                .publish(
                        new CommunicationIdRemovedEvent(removedId));
    }
}