package net.tomofiles.skysign.vehicle.domain.vehicle;

public class CarbonCopyVehicleService {
    public static void copy(
        Generator generator,
        VehicleRepository repository,
        VehicleId originalId,
        VehicleId newId
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

    }
}