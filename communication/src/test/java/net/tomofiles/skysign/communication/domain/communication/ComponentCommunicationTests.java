package net.tomofiles.skysign.communication.domain.communication;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.time.LocalDateTime;
import java.util.Arrays;

import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;

public class ComponentCommunicationTests {
    
    /**
     * DTOからCommunicationエンティティを組み立てる。
     */
    @Test
    public void assembleIntoCommunicationTest() {
        CommunicationId id = CommunicationId.newId();
        MissionId missionId = new MissionId("new mission id");
        Version version = new Version(1);

        double latitude = 0.0d;
        double longitude = 1.0d;
        double altitude = 2.0d;
        double relativeAltitude = 3.0d;
        double speed = 4.0d;
        boolean armed = true;
        String flightMode = "INFLIGHT";
        double orientationX = 5.0d;
        double orientationY = 6.0d;
        double orientationZ = 7.0d;
        double orientationW = 8.0d;

        CommandId commandId1 = CommandId.newId();
        CommandType type1 = CommandType.ARM;
        LocalDateTime time1 = LocalDateTime.of(2020, 07, 22, 10, 30, 25);
        CommandId commandId2 = CommandId.newId();
        CommandType type2 = CommandType.DISARM;
        LocalDateTime time2 = LocalDateTime.of(2020, 07, 22, 10, 30, 30);
        CommandId commandId3 = CommandId.newId();
        CommandType type3 = CommandType.UPLOAD;
        LocalDateTime time3 = LocalDateTime.of(2020, 07, 22, 10, 30, 45);

        Communication communication = CommunicationFactory.assembleFrom(
                new CommunicationComponentDto(
                        id.getId(),
                        missionId.getId(),
                        version.getVersion(),
                        new TelemetryComponentDto(
                                latitude,
                                longitude,
                                altitude,
                                relativeAltitude,
                                speed,
                                armed,
                                flightMode,
                                orientationX,
                                orientationY,
                                orientationZ,
                                orientationW),
                        Arrays.asList(new CommandComponentDto[] {
                                new CommandComponentDto(
                                        commandId1.getId(),
                                        type1.name(),
                                        time1),
                                new CommandComponentDto(
                                        commandId2.getId(),
                                        type2.name(),
                                        time2),
                                new CommandComponentDto(
                                        commandId3.getId(),
                                        type3.name(),
                                        time3)
                        })
                )
        );

        Telemetry after = Telemetry.newInstance()
                .setPosition(latitude, longitude, altitude, relativeAltitude, speed)
                .setArmed(armed)
                .setFlightMode(flightMode)
                .setOrientation(orientationX, orientationY, orientationZ, orientationW);

        assertEquals(communication.getId(), id);
        assertEquals(communication.getMissionId(), missionId);
        assertEquals(communication.getCommands().get(0), new Command(commandId1, type1, time1));
        assertEquals(communication.getCommands().get(1), new Command(commandId2, type2, time2));
        assertEquals(communication.getCommands().get(2), new Command(commandId3, type3, time3));
        assertEquals(communication.getTelemetry(), after);
        assertEquals(communication.getVersion(), version);
    }

    /**
     * CommunicationエンティティからDTOに分解する。
     */
    @Test
    public void takeApartCommunicationTest() {
        CommunicationId id = CommunicationId.newId();
        MissionId missionId = new MissionId("new mission id");
        Version version = new Version(1);

        double latitude = 0.0d;
        double longitude = 1.0d;
        double altitude = 2.0d;
        double relativeAltitude = 3.0d;
        double speed = 4.0d;
        boolean armed = true;
        String flightMode = "INFLIGHT";
        double orientationX = 5.0d;
        double orientationY = 6.0d;
        double orientationZ = 7.0d;
        double orientationW = 8.0d;

        CommandId commandId1 = CommandId.newId();
        CommandType type1 = CommandType.ARM;
        LocalDateTime time1 = LocalDateTime.of(2020, 07, 22, 10, 30, 25);
        CommandId commandId2 = CommandId.newId();
        CommandType type2 = CommandType.DISARM;
        LocalDateTime time2 = LocalDateTime.of(2020, 07, 22, 10, 30, 30);
        CommandId commandId3 = CommandId.newId();
        CommandType type3 = CommandType.UPLOAD;
        LocalDateTime time3 = LocalDateTime.of(2020, 07, 22, 10, 30, 45);

        Communication communication = CommunicationFactory.newInstance(id);
        communication.setMissionId(missionId);
        communication.setTelemetry(
                Telemetry.newInstance()
                    .setPosition(latitude, longitude, altitude, relativeAltitude, speed)
                    .setArmed(armed)
                    .setFlightMode(flightMode)
                    .setOrientation(orientationX, orientationY, orientationZ, orientationW)
        );
        communication.getCommands().add(new Command(commandId1, type1, time1));
        communication.getCommands().add(new Command(commandId2, type2, time2));
        communication.getCommands().add(new Command(commandId3, type3, time3));

        CommunicationComponentDto dto = CommunicationFactory.takeApart(communication);

        assertEquals(dto.getId(), id.getId());
        assertEquals(dto.getMissionId(), missionId.getId());
        assertEquals(dto.getCommands().get(0).getId(), commandId1.getId());
        assertEquals(dto.getCommands().get(0).getType(), type1.name());
        assertEquals(dto.getCommands().get(0).getTime(), time1);
        assertEquals(dto.getCommands().get(1).getId(), commandId2.getId());
        assertEquals(dto.getCommands().get(1).getType(), type2.name());
        assertEquals(dto.getCommands().get(1).getTime(), time2);
        assertEquals(dto.getCommands().get(2).getId(), commandId3.getId());
        assertEquals(dto.getCommands().get(2).getType(), type3.name());
        assertEquals(dto.getCommands().get(2).getTime(), time3);
        assertEquals(dto.getTelemetry().getLatitude(), latitude);
        assertEquals(dto.getTelemetry().getLongitude(), longitude);
        assertEquals(dto.getTelemetry().getAltitude(), altitude);
        assertEquals(dto.getTelemetry().getRelativeAltitude(), relativeAltitude);
        assertEquals(dto.getTelemetry().getSpeed(), speed);
        assertEquals(dto.getTelemetry().isArmed(), armed);
        assertEquals(dto.getTelemetry().getFlightMode(), flightMode);
        assertEquals(dto.getTelemetry().getOriX(), orientationX);
        assertEquals(dto.getTelemetry().getOriY(), orientationY);
        assertEquals(dto.getTelemetry().getOriZ(), orientationZ);
        assertEquals(dto.getTelemetry().getOriW(), orientationW);
        assertEquals(dto.getVersion(), version.getVersion());
    }
}