package net.tomofiles.skysign.communication.service;

import java.util.List;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.vehicle.Generator;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.communication.event.Publisher;
import net.tomofiles.skysign.communication.service.dpo.CreateVehicleRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.CreateVehicleResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.DeleteVehicleRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.DeleteVehicleResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.GetVehicleRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.GetVehicleResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.ListVehiclesResponsesDpo;
import net.tomofiles.skysign.communication.service.dpo.UpdateVehicleRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.UpdateVehicleResponseDpo;

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

        vehicle.removeCommId();

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