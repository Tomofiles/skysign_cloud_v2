package net.tomofiles.skysign.communication.usecase;

import java.util.List;
import java.util.NoSuchElementException;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.communication.event.Publisher;
import net.tomofiles.skysign.communication.usecase.dto.VehicleDto;

@Component
@AllArgsConstructor
public class ManageVehicleService {

    private final VehicleRepository vehicleRepository;
    private final Publisher publisher;

    @Transactional
    public String createVehicle(String name, String commId) {
        VehicleId id = VehicleId.newId();
        Vehicle vehicle = VehicleFactory.newInstance(id);

        vehicle.setPublisher(this.publisher);

        vehicle.nameVehicle(name);
        vehicle.giveCommId(new CommunicationId(commId));

        this.vehicleRepository.save(vehicle);

        return id.getId();
    }

    @Transactional
    public void updateVehicle(String vehicleId, String name, String commId) {
        VehicleId id = new VehicleId(vehicleId);
        Vehicle vehicle = this.vehicleRepository.getById(id);

        if (vehicle == null) {
            throw new NoSuchElementException("vehicle-idに合致するVehicleが存在しません。");
        }

        vehicle.setPublisher(this.publisher);

        vehicle.nameVehicle(name);
        vehicle.giveCommId(new CommunicationId(commId));

        this.vehicleRepository.save(vehicle);
    }

    @Transactional
    public void deleteVehicle(String vehicleId) {
        VehicleId id = new VehicleId(vehicleId);
        Vehicle vehicle = this.vehicleRepository.getById(id);

        if (vehicle == null) {
            throw new NoSuchElementException("vehicle-idに合致するVehicleが存在しません。");
        }

        vehicle.setPublisher(this.publisher);

        this.vehicleRepository.remove(id, vehicle.getVersion());
    }

    @Transactional
    public VehicleDto getVehicle(String vehicleId) {
        VehicleId id = new VehicleId(vehicleId);
        Vehicle vehicle = this.vehicleRepository.getById(id);

        if (vehicle == null) {
            throw new NoSuchElementException("vehicle-idに合致するVehicleが存在しません。");
        }

        VehicleDto dto = new VehicleDto();
        dto.setId(vehicle.getId().getId());
        dto.setName(vehicle.getVehicleName());
        dto.setCommId(vehicle.getCommId().getId());

        return dto;
    }

    @Transactional
    public List<VehicleDto> getAllVehicle() {
        List<Vehicle> vehicles = this.vehicleRepository.getAll();

        List<VehicleDto> dtos = vehicles.stream()
                .map(vehicle -> {
                    VehicleDto dto = new VehicleDto();
                    dto.setId(vehicle.getId().getId());
                    dto.setName(vehicle.getVehicleName());
                    dto.setCommId(vehicle.getCommId().getId());
                    return dto;
                })
                .collect(Collectors.toList());

        return dtos;
    }
}