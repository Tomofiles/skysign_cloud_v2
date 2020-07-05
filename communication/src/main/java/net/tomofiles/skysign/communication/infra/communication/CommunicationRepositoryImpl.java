package net.tomofiles.skysign.communication.infra.communication;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;
import net.tomofiles.skysign.communication.infra.common.DeleteCondition;

@Component
public class CommunicationRepositoryImpl implements CommunicationRepository {

    @Autowired
    private CommunicationMapper communicationMapper;

    @Autowired
    private TelemetryMapper telemetryMapper;

    @Autowired
    private CommandMapper commandMapper;

    @Override
    public void save(Communication comm) {
        boolean isCreate = false;
        TelemetryRecord telemetry = null;
        List<CommandRecord> commandsInDB = new ArrayList<>();

        CommunicationComponentDto componentDto = CommunicationFactory.takeApart(comm); 

        CommunicationRecord communication = this.communicationMapper.find(componentDto.getId());

        if (communication == null) {
            communication = new CommunicationRecord();
            communication.setId(componentDto.getId());
            isCreate = true;

            telemetry = new TelemetryRecord();
            telemetry.setCommId(componentDto.getId());
        } else {
            telemetry = this.telemetryMapper.find(componentDto.getId());
            commandsInDB.addAll(this.commandMapper.findByCommId(componentDto.getId()));
        }

        communication.setMissionId(componentDto.getMissionId());
        communication.setVersion(componentDto.getVersion());

        telemetry.setLatitude(componentDto.getTelemetry().getLatitude());
        telemetry.setLongitude(componentDto.getTelemetry().getLongitude());
        telemetry.setAltitude(componentDto.getTelemetry().getAltitude());
        telemetry.setRelativeAltitude(componentDto.getTelemetry().getRelativeAltitude());
        telemetry.setSpeed(componentDto.getTelemetry().getSpeed());
        telemetry.setArmed(componentDto.getTelemetry().isArmed());
        telemetry.setFlightMode(componentDto.getTelemetry().getFlightMode());
        telemetry.setOriX(componentDto.getTelemetry().getOriX());
        telemetry.setOriY(componentDto.getTelemetry().getOriY());
        telemetry.setOriZ(componentDto.getTelemetry().getOriZ());
        telemetry.setOriW(componentDto.getTelemetry().getOriW());

        List<CommandRecord> commands = componentDto.getCommands().stream()
                .map(c -> {
                        return new CommandRecord(c.getId(), componentDto.getId(), c.getType(), c.getTime());
                })
                .collect(Collectors.toList());

        if (isCreate) {
            this.communicationMapper.create(communication);
            this.telemetryMapper.create(telemetry);
            commands.stream()
                    .forEach(this.commandMapper::create);
        } else {
            this.communicationMapper.update(communication);
            this.telemetryMapper.update(telemetry);

            commands.stream()
                    .filter(c -> !commandsInDB.contains(c))
                    .forEach(this.commandMapper::create);
            commandsInDB.stream()
                    .filter(c -> !commands.contains(c))
                    .map(CommandRecord::getId)
                    .forEach(this.commandMapper::delete);
        }
    }

    @Override
    public void remove(CommunicationId id, Version version) {
        DeleteCondition condition = new DeleteCondition();
        
        condition.setId(id.getId());
        condition.setVersion(version.getVersion());
        
        this.communicationMapper.delete(condition);
        this.telemetryMapper.delete(id.getId());
        this.commandMapper.deleteByCommId(id.getId());
    }

    @Override
    public Communication getById(CommunicationId id) {
        CommunicationRecord communication = this.communicationMapper.find(id.getId());

        if (communication == null) {
            return null;
        }

        TelemetryRecord telemetry = this.telemetryMapper.find(id.getId());
        List<CommandRecord> commands = this.commandMapper.findByCommId(communication.getId());

        return CommunicationFactory.assembleFrom(
                new CommunicationComponentDto(
                        id.getId(),
                        communication.getMissionId(),
                        communication.getVersion(),
                        new TelemetryComponentDto(
                                telemetry.getLatitude(),
                                telemetry.getLongitude(),
                                telemetry.getAltitude(),
                                telemetry.getRelativeAltitude(),
                                telemetry.getSpeed(),
                                telemetry.isArmed(),
                                telemetry.getFlightMode(),
                                telemetry.getOriX(),
                                telemetry.getOriY(),
                                telemetry.getOriZ(),
                                telemetry.getOriW()),
                        commands.stream()
                                .map(c -> new CommandComponentDto(
                                    c.getId(),
                                    c.getType(),
                                    c.getTime()
                                ))
                                .collect(Collectors.toList())));
    }
}