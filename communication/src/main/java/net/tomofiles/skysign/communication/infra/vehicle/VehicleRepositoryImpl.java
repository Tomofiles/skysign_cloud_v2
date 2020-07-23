package net.tomofiles.skysign.communication.infra.vehicle;

import java.util.List;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.communication.domain.vehicle.Version;
import net.tomofiles.skysign.communication.infra.common.DeleteCondition;

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
        record.setNewVersion(vehicle.getNewVersion().getVersion());

        if (isCreate) {
            this.vehicleMapper.create(record);
        } else {
            this.vehicleMapper.update(record);
        }
    }

    @Override
    public void remove(VehicleId id, Version version) {
        DeleteCondition condition = new DeleteCondition();

        condition.setId(id.getId());
        condition.setVersion(version.getVersion());
        
        this.vehicleMapper.delete(condition);
    }

    @Override
    public Vehicle getById(VehicleId id) {
        VehicleRecord record = this.vehicleMapper.find(id.getId());

        if (record == null) {
            return null;
        }

        return VehicleFactory.rebuild(id, record.getName(), record.getCommId(), record.getVersion());
    }

    @Override
    public List<Vehicle> getAll() {
        List<VehicleRecord> records = this.vehicleMapper.findAll();

        List<Vehicle> results = records.stream()
                .map(record -> VehicleFactory.rebuild(new VehicleId(record.getId()), record.getName(), record.getCommId(), record.getVersion()))
                .collect(Collectors.toList());

        return results;
    }
}