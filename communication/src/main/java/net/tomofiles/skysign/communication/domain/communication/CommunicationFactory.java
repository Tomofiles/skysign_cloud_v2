package net.tomofiles.skysign.communication.domain.communication;

import java.util.stream.Collectors;

import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;

public class CommunicationFactory {

    public static Communication newInstance(CommunicationId communicationId, Generator generator) {
        Communication communication = new Communication(communicationId, generator);
        communication.setTelemetry(Telemetry.newInstance());
        return communication;
    }

    public static Communication assembleFrom(CommunicationComponentDto componentDto, Generator generator) {
        Communication communication = new Communication(
                new CommunicationId(componentDto.getId()),
                generator
        );
        communication.setMissionId(new MissionId(componentDto.getMissionId()));
        communication.setTelemetry(Telemetry.newInstance()
                .setPosition(
                        componentDto.getTelemetry().getLatitude(),
                        componentDto.getTelemetry().getLongitude(),
                        componentDto.getTelemetry().getAltitude(),
                        componentDto.getTelemetry().getRelativeAltitude(),
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
                                    CommandType.valueOf(c.getType()),
                                    c.getTime());
                        })
                        .collect(Collectors.toList())
        );
        return communication;
    }

    public static CommunicationComponentDto takeApart(Communication communication) {
        return  new CommunicationComponentDto(
                communication.getId().getId(),
                communication.getMissionId() == null ? null : communication.getMissionId().getId(),
                new TelemetryComponentDto(
                        communication.getTelemetry().getPosition().getLatitude(),
                        communication.getTelemetry().getPosition().getLongitude(),
                        communication.getTelemetry().getPosition().getAltitude(),
                        communication.getTelemetry().getPosition().getRelativeAltitude(),
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
                            c.getType().toString(),
                            c.getTime()
                        ))
                        .collect(Collectors.toList()));
    }
}