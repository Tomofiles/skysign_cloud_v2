package net.tomofiles.skysign.communication.usecase;

import java.util.List;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.vehicle.Generator;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.communication.event.Publisher;
import net.tomofiles.skysign.communication.usecase.dpo.CreateVehicleRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.CreateVehicleResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.DeleteVehicleRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.DeleteVehicleResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.GetVehicleRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.GetVehicleResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.ListVehiclesResponsesDpo;
import net.tomofiles.skysign.communication.usecase.dpo.UpdateVehicleRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.UpdateVehicleResponseDpo;

@Component
@AllArgsConstructor
public class ManageVehicleService {

    private final VehicleRepository vehicleRepository;
    private final Publisher publisher;
    private final Generator generator;

    @Transactional
    public void createVehicle(CreateVehicleRequestDpo requestDpo, CreateVehicleResponseDpo responseDpo) {
        Vehicle vehicle = VehicleFactory.newInstance(this.generator);

        vehicle.setPublisher(this.publisher);

        vehicle.nameVehicle(requestDpo.getVehicleName());
        vehicle.giveCommId(requestDpo.getCommId());

        this.vehicleRepository.save(vehicle);

        responseDpo.setVehicle(vehicle);
    }

    @Transactional
    public void updateVehicle(UpdateVehicleRequestDpo requestDpo, UpdateVehicleResponseDpo responseDpo) {
        Vehicle vehicle = this.vehicleRepository.getById(requestDpo.getVehicleId());

        if (vehicle == null) {
            return;
        }

        vehicle.setPublisher(this.publisher);

        vehicle.nameVehicle(requestDpo.getVehicleName());
        vehicle.giveCommId(requestDpo.getCommId());

        this.vehicleRepository.save(vehicle);

        responseDpo.setVehicle(vehicle);
    }

    @Transactional
    public void deleteVehicle(DeleteVehicleRequestDpo requestDpo, DeleteVehicleResponseDpo responseDpo) {
        Vehicle vehicle = this.vehicleRepository.getById(requestDpo.getVehicleId());

        if (vehicle == null) {
            return;
        }

        vehicle.setPublisher(this.publisher);

        this.vehicleRepository.remove(vehicle.getId(), vehicle.getVersion());

        responseDpo.setVehicle(vehicle);
    }

    @Transactional
    public void getVehicle(GetVehicleRequestDpo requestDpo, GetVehicleResponseDpo responseDpo) {
        Vehicle vehicle = this.vehicleRepository.getById(requestDpo.getVehicleId());

        responseDpo.setVehicle(vehicle);
    }

    @Transactional
    public void listVehicles(ListVehiclesResponsesDpo responsesDpo) {
        List<Vehicle> vehicles = this.vehicleRepository.getAll();

        responsesDpo.setVehicles(vehicles);
    }
}