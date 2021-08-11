package net.tomofiles.skysign.communication.infra.communication;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.Generator;
import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.UploadMissionComponentDto;

@Component
@RequiredArgsConstructor
public class CommunicationRepositoryImpl implements CommunicationRepository {

    private final CommunicationMapper communicationMapper;
    private final TelemetryMapper telemetryMapper;
    private final CommandMapper commandMapper;
    private final UploadMissionMapper uploadMissionMapper;
    private final Generator generator;

    @Override
    public void save(Communication comm) {
        boolean isCreate = false;
        TelemetryRecord telemetry = null;
        List<CommandRecord> commandsInDB = new ArrayList<>();
        List<UploadMissionRecord> uploadMissionsInDB = new ArrayList<>();

        CommunicationComponentDto componentDto = CommunicationFactory.takeApart(comm); 

        CommunicationRecord communication = this.communicationMapper.find(componentDto.getId());

        if (communication == null) {
            communication = new CommunicationRecord();
            communication.setId(componentDto.getId());
            isCreate = true;

            telemetry = new TelemetryRecord();
            telemetry.setCommunicationId(componentDto.getId());
        } else {
            telemetry = this.telemetryMapper.find(componentDto.getId());
            commandsInDB.addAll(this.commandMapper.findByCommunicationId(componentDto.getId()));
            uploadMissionsInDB.addAll(this.uploadMissionMapper.findByCommunicationId(componentDto.getId()));
        }

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

        List<UploadMissionRecord> uploadMissions = componentDto.getUploadMissions().stream()
                .map(um -> {
                        return new UploadMissionRecord(um.getId(), componentDto.getId(), um.getMissionId());
                })
                .collect(Collectors.toList());

        if (isCreate) {
            this.communicationMapper.create(communication);
            this.telemetryMapper.create(telemetry);
            commands.stream()
                    .forEach(this.commandMapper::create);
            uploadMissions.stream()
                    .forEach(this.uploadMissionMapper::create);
        } else {
            this.telemetryMapper.update(telemetry);

            commands.stream()
                    .filter(c -> !commandsInDB.contains(c))
                    .forEach(this.commandMapper::create);
            commandsInDB.stream()
                    .filter(c -> !commands.contains(c))
                    .map(CommandRecord::getId)
                    .forEach(this.commandMapper::delete);

            uploadMissions.stream()
                    .filter(um -> !uploadMissionsInDB.contains(um))
                    .forEach(this.uploadMissionMapper::create);
            uploadMissionsInDB.stream()
                    .filter(um -> !uploadMissions.contains(um))
                    .map(UploadMissionRecord::getId)
                    .forEach(this.uploadMissionMapper::delete);
        }
    }

    @Override
    public void remove(CommunicationId id) {
        this.communicationMapper.delete(id.getId());
        this.telemetryMapper.delete(id.getId());
        this.commandMapper.deleteByCommunicationId(id.getId());
    }

    @Override
    public Communication getById(CommunicationId id) {
        CommunicationRecord communication = this.communicationMapper.find(id.getId());

        if (communication == null) {
            return null;
        }

        TelemetryRecord telemetry = this.telemetryMapper.find(id.getId());
        List<CommandRecord> commands = this.commandMapper.findByCommunicationId(communication.getId());
        List<UploadMissionRecord> uploadMissions = this.uploadMissionMapper.findByCommunicationId(communication.getId());

        return CommunicationFactory.assembleFrom(
                new CommunicationComponentDto(
                        id.getId(),
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
                                .collect(Collectors.toList()),
                        uploadMissions.stream()
                                .map(c -> new UploadMissionComponentDto(
                                    c.getId(),
                                    c.getMissionId()
                                ))
                                .collect(Collectors.toList())
                    ),
                generator
        );
    }

    @Override
    public List<Communication> getAll() {
        List<CommunicationRecord> commRecords = this.communicationMapper.findAll();

        if (commRecords.isEmpty()) {
            return Collections.emptyList();
        }

        List<Communication> communications = new ArrayList<>();
        for (CommunicationRecord commRecord : commRecords) {

            TelemetryRecord telemetry = this.telemetryMapper.find(commRecord.getId());
            List<CommandRecord> commands = this.commandMapper.findByCommunicationId(commRecord.getId());
            List<UploadMissionRecord> uploadMissions = this.uploadMissionMapper.findByCommunicationId(commRecord.getId());
    
            Communication communication = CommunicationFactory.assembleFrom(
                    new CommunicationComponentDto(
                            commRecord.getId(),
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
                                    .collect(Collectors.toList()),
                            uploadMissions.stream()
                                    .map(c -> new UploadMissionComponentDto(
                                        c.getId(),
                                        c.getMissionId()
                                    ))
                                    .collect(Collectors.toList())
                        ),
                    generator
            );
            communications.add(communication);
        }

        return communications;
    }
}