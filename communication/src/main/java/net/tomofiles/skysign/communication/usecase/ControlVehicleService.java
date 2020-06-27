package net.tomofiles.skysign.communication.usecase;

import java.util.NoSuchElementException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.MissionId;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.communication.usecase.dto.ControlCommandType;
import net.tomofiles.skysign.communication.usecase.dto.TelemetryDto;

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

    @Transactional
    public void cancel(String vehicleId) {
        VehicleId id = new VehicleId(vehicleId);
        Vehicle vehicle = this.vehicleRepository.getById(id);

        if (vehicle == null) {
            throw new NoSuchElementException("vehicle-idに合致するVehicleが存在しません。");
        }

        Communication communication = this.communicationRepository.getById(vehicle.getCommId());

        if (communication == null) {
            throw new IllegalStateException("vehicleにcommunication-idが設定されていません。");
        }

        communication.cancel();

        this.communicationRepository.save(communication);
    }

    @Transactional
    public void pushCommand(String vehicleId, ControlCommandType commandType) {
        VehicleId id = new VehicleId(vehicleId);
        Vehicle vehicle = this.vehicleRepository.getById(id);

        if (vehicle == null) {
            throw new NoSuchElementException("vehicle-idに合致するVehicleが存在しません。");
        }

        Communication communication = this.communicationRepository.getById(vehicle.getCommId());

        if (communication == null) {
            throw new IllegalStateException("vehicleにcommunication-idが設定されていません。");
        }

        communication.pushCommand(commandType.getType());

        this.communicationRepository.save(communication);
    }

    @Transactional
    public TelemetryDto pullTelemetry(String vehicleId) {
        VehicleId id = new VehicleId(vehicleId);
        Vehicle vehicle = this.vehicleRepository.getById(id);

        if (vehicle == null) {
            throw new NoSuchElementException("vehicle-idに合致するVehicleが存在しません。");
        }

        Communication communication = this.communicationRepository.getById(vehicle.getCommId());

        if (communication == null) {
            throw new IllegalStateException("vehicleにcommunication-idが設定されていません。");
        }

        TelemetrySnapshot telemetry = communication.pullTelemetry();

        TelemetryDto dto = new TelemetryDto();
        dto.setLatitude(telemetry.getLatitude());
        dto.setLongitude(telemetry.getLongitude());
        dto.setAltitude(telemetry.getAltitude());
        dto.setSpeed(telemetry.getSpeed());
        dto.setArmed(telemetry.isArmed());
        dto.setFlightMode(telemetry.getFlightMode());
        dto.setOrientationX(telemetry.getX());
        dto.setOrientationY(telemetry.getY());
        dto.setOrientationZ(telemetry.getZ());
        dto.setOrientationW(telemetry.getW());

        return dto;
    }
}