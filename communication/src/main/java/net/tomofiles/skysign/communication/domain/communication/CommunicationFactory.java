package net.tomofiles.skysign.communication.domain.communication;

import java.util.ArrayList;
import java.util.stream.Collectors;

import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;

public class CommunicationFactory {

    public static Communication newInstance(CommunicationId id) {
        Communication communication = new Communication(id, new ArrayList<>());
        communication.setTelemetry(Telemetry.newInstance());
        communication.setVersion(new Version(1));
        return communication;
    }

    public static Communication assembleFrom(CommunicationComponentDto componentDto) {
        Communication communication = new Communication(new CommunicationId(componentDto.getId()), new ArrayList<>());
        communication.setMissionId(new MissionId(componentDto.getMissionId()));
        communication.setVersion(new Version(componentDto.getVersion()));
        communication.setTelemetry(Telemetry.newInstance()
                .setPosition(
                        componentDto.getTelemetry().getLatitude(),
                        componentDto.getTelemetry().getLongitude(),
                        componentDto.getTelemetry().getAltitude(),
                        componentDto.getTelemetry().getSpeed())
                .setArmed(componentDto.getTelemetry().isArmed())
                .setFlightMode(componentDto.getTelemetry().getFlightMode())
                .setOrientation(
                        componentDto.getTelemetry().getOriX(),
                        componentDto.getTelemetry().getOriY(),
                        componentDto.getTelemetry().getOriZ(),
                        componentDto.getTelemetry().getOriW()
                ));
        communication.getCommands().addAll(
                componentDto.getCommands().stream()
                        .map(c -> {
                                return new Command(
                                    new CommandId(c.getId()),
                                    CommandType.valueOf(c.getType()));
                        })
                        .collect(Collectors.toList())
        );
        return communication;
    }

    public static CommunicationComponentDto takeApart(Communication communication) {
        return  new CommunicationComponentDto(
                communication.getId().getId(),
                communication.getMissionId() == null ? null : communication.getMissionId().getId(),
                communication.getVersion().getVersion(),
                new TelemetryComponentDto(
                        communication.getTelemetry().getPosition().getLatitude(),
                        communication.getTelemetry().getPosition().getLongitude(),
                        communication.getTelemetry().getPosition().getAltitude(),
                        communication.getTelemetry().getSpeed(),
                        communication.getTelemetry().isArmed(),
                        communication.getTelemetry().getFlightMode(),
                        communication.getTelemetry().getOrientation().getX(),
                        communication.getTelemetry().getOrientation().getY(),
                        communication.getTelemetry().getOrientation().getZ(),
                        communication.getTelemetry().getOrientation().getW()),
                communication.getCommands().stream()
                        .map(c -> new CommandComponentDto(
                            c.getId().getId(),
                            c.getType().toString()
                        ))
                        .collect(Collectors.toList()));
    }
}