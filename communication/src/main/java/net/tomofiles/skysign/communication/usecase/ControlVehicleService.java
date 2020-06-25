package net.tomofiles.skysign.communication.usecase;

import java.util.NoSuchElementException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.MissionId;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;

@Component
public class ControlVehicleService {

    @Autowired
    private VehicleRepository vehicleRepository;

    @Autowired
    private CommunicationRepository communicationRepository;

    @Transactional
    public void standBy(String vehicleId, String missionId) {
        VehicleId id = new VehicleId(vehicleId);
        Vehicle vehicle = this.vehicleRepository.getById(id);

        if (vehicle == null) {
            throw new NoSuchElementException("vehicle-idに合致するVehicleが存在しません。");
        }

        Communication communication = this.communicationRepository.getById(vehicle.getCommId());

        if (communication == null) {
            throw new IllegalStateException("vehicleにcommunication-idが設定されていません。");
        }

        communication.standBy(new MissionId(missionId));

        this.communicationRepository.save(communication);
    }
}