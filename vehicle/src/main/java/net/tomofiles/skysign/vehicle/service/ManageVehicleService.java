package net.tomofiles.skysign.vehicle.service;

import java.util.List;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.vehicle.domain.vehicle.CarbonCopyVehicleService;
import net.tomofiles.skysign.vehicle.domain.vehicle.Generator;
import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.vehicle.event.Publisher;
import net.tomofiles.skysign.vehicle.service.dpo.CarbonCopyVehicleRequestDpo;
import net.tomofiles.skysign.vehicle.service.dpo.CreateVehicleRequestDpo;
import net.tomofiles.skysign.vehicle.service.dpo.CreateVehicleResponseDpo;
import net.tomofiles.skysign.vehicle.service.dpo.DeleteVehicleRequestDpo;
import net.tomofiles.skysign.vehicle.service.dpo.DeleteVehicleResponseDpo;
import net.tomofiles.skysign.vehicle.service.dpo.GetVehicleRequestDpo;
import net.tomofiles.skysign.vehicle.service.dpo.GetVehicleResponseDpo;
import net.tomofiles.skysign.vehicle.service.dpo.ListVehiclesResponsesDpo;
import net.tomofiles.skysign.vehicle.service.dpo.UpdateVehicleRequestDpo;
import net.tomofiles.skysign.vehicle.service.dpo.UpdateVehicleResponseDpo;

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
        List<Vehicle> vehicles = this.vehicleRepository.getAllOriginal();

        responsesDpo.setVehicles(vehicles);
    }

    @Transactional
    public void carbonCopyVehicle(CarbonCopyVehicleRequestDpo requestDpo) {
        CarbonCopyVehicleService.copy(
            this.generator,
            this.publisher,
            this.vehicleRepository, 
            requestDpo.getOriginalId(), 
            requestDpo.getNewId(),
            requestDpo.getFlightplanId());
    }
}