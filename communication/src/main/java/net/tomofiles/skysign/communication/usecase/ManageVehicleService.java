package net.tomofiles.skysign.communication.usecase;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.communication.event.Publisher;

@Component
public class ManageVehicleService {

    @Autowired
    private VehicleRepository vehicleRepository;

    @Autowired
    private Publisher publisher;

    @Transactional
    public void createVehicle(String name, String commId) {
        VehicleId id = VehicleId.newId();
        Vehicle vehicle = VehicleFactory.newInstance(id);

        vehicle.setPublisher(this.publisher);

        vehicle.namedVehicle(name);
        vehicle.giveCommId(new CommunicationId(commId));

        this.vehicleRepository.save(vehicle);
    }
}