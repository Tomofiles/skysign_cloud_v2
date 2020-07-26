package net.tomofiles.skysign.communication.usecase;

import java.util.List;
import java.util.NoSuchElementException;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.MissionId;
import net.tomofiles.skysign.communication.usecase.dto.ControlCommandDto;
import net.tomofiles.skysign.communication.usecase.dto.ControlCommandType;
import net.tomofiles.skysign.communication.usecase.dto.TelemetryDto;

@Component
@AllArgsConstructor
public class CommunicateEdgeService {

    private final CommunicationRepository communicationRepository;

    @Transactional
    public List<String> pushTelemetry(String commId, TelemetryDto telemetry) {
        CommunicationId id = new CommunicationId(commId);
        Communication communication = this.communicationRepository.getById(id);

        if (communication == null) {
            throw new NoSuchElementException("communication-idに合致するCommunicationが存在しません。");
        }

        communication.pushTelemetry(
                telemetry.getLatitude(),
                telemetry.getLongitude(),
                telemetry.getAltitude(),
                telemetry.getRelativeAltitude(),
                telemetry.getSpeed(),
                telemetry.isArmed(),
                telemetry.getFlightMode(),
                telemetry.getOrientationX(),
                telemetry.getOrientationY(),
                telemetry.getOrientationZ(),
                telemetry.getOrientationW());
        
        List<String> commandIds = communication.getCommandIds().stream()
                .map(CommandId::getId)
                .collect(Collectors.toList());

        this.communicationRepository.save(communication);

        return commandIds;
    }

    @Transactional
    public ControlCommandDto pullCommand(String commId, String commandId) {
        CommunicationId id = new CommunicationId(commId);
        CommandId cid = new CommandId(commandId);

        Communication communication = this.communicationRepository.getById(id);

        if (communication == null) {
            throw new NoSuchElementException("communication-idに合致するCommunicationが存在しません。");
        }

        CommandType type = communication.pullCommandById(cid);
        
        if (type == null) {
            throw new NoSuchElementException("command-idに合致するCommandが存在しません。");
        }

        this.communicationRepository.save(communication);

        if (type == CommandType.UPLOAD) {
            MissionId missionId = communication.getMissionId();
            return new ControlCommandDto(ControlCommandType.valueOf(type), missionId.getId());
        }
        return new ControlCommandDto(ControlCommandType.valueOf(type), null);
    }
}