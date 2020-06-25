package net.tomofiles.skysign.communication.infra.vehicle;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;

@Component
public class VehicleRepositoryImpl implements VehicleRepository {

    @Autowired
    private VehicleMapper vehicleMapper;

    @Override
    public void save(Vehicle vehicle) {
        boolean isCreate = false;

        VehicleRecord record = this.vehicleMapper.find(vehicle.getId().getId());

        if (record == null) {
            record = new VehicleRecord();
            record.setId(vehicle.getId().getId());
            isCreate = true;
        }

        record.setName(vehicle.getVehicleName());
        record.setCommId(vehicle.getCommId().getId());
        record.setVersion(vehicle.getVersion().getVersion());

        if (isCreate) {
            this.vehicleMapper.create(record);
        } else {
            this.vehicleMapper.update(record);
        }
    }

    @Override
    public Vehicle getById(VehicleId id) {
        VehicleRecord record = this.vehicleMapper.find(id.getId());

        if (record == null) {
            return null;
        }

        return VehicleFactory.rebuild(
            id, 
            record.getName(), 
            record.getCommId(), 
            record.getVersion());
    }
}