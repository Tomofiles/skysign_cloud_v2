package net.tomofiles.skysign.vehicle.domain.vehicle;

import net.tomofiles.skysign.vehicle.event.Publisher;

public class CarbonCopyVehicleService {
    public static void copy(
        Generator generator,
        Publisher publisher,
        VehicleRepository repository,
        VehicleId originalId,
        VehicleId newId,
        FlightplanId flightplanId
    ) {
        Vehicle newVehicle = repository.getById(newId);

        if (newVehicle != null) {
            return;
        }

        Vehicle original = repository.getById(originalId);

        if (original == null) {
            return;
        }

        Vehicle copy = VehicleFactory.copy(original, newId, generator);

        repository.save(copy);

        publisher.publish(
            new CopiedVehicleCreatedEvent(
                copy.getId(),
                copy.getCommunicationId(),
                flightplanId
            )
        );
    }
}