package net.tomofiles.skysign.communication.domain.communication;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;

public class CommunicationObjectMother {
    
    /**
     * テスト用Communicationエンティティを生成する。
     */
    public static Communication newNormalCommunication(
            CommunicationId communicationId,
            VehicleId vehicleId,
            boolean controlled,
            Generator generator) {
        Communication communication = CommunicationFactory.newInstance(communicationId, vehicleId, generator);
        communication.setControlled(controlled);
        communication.setTelemetry(newNormalTelemetry());
        return communication;
    }

    /**
     * 1件のCommandを持つテスト用Communicationエンティティを生成する。
     */
    public static Communication newSingleCommandCommunication(
            CommunicationId communicationId,
            VehicleId vehicleId,
            boolean controlled,
            Generator generator,
            Generator generatorCommand,
            Generator generatorUploadMission) {
        Communication communication = CommunicationFactory.newInstance(communicationId, vehicleId, generator);
        communication.setControlled(controlled);
        communication.getCommands().addAll(newSingleCommands(generatorCommand, CommandType.ARM));
        communication.getUploadMissions().addAll(newSingleUploadMissions(generatorUploadMission));
        communication.setTelemetry(newNormalTelemetry());
        return communication;
    }

    /**
     * 複数件のCommandリストを持つテスト用Communicationエンティティを生成する。
     */
    public static Communication newSeveralCommandsCommunication(
            CommunicationId communicationId,
            VehicleId vehicleId,
            boolean controlled,
            Generator generator,
            Generator generatorCommand,
            Generator generatorUploadMission) {
        Communication communication = CommunicationFactory.newInstance(communicationId, vehicleId, generator);
        communication.setControlled(controlled);
        communication.getCommands().addAll(newSeveralCommands(generatorCommand));
        communication.getUploadMissions().addAll(newSeveralUploadMissions(generatorUploadMission));
        communication.setTelemetry(newNormalTelemetry());
        return communication;
    }

    /**
     * テスト用Telemetryオブジェクトを生成する。
     */
    public static Telemetry newNormalTelemetry() {
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

        return Telemetry.newInstance()
                .setPosition(latitude, longitude, altitude, relativeAltitude, speed)
                .setArmed(armed)
                .setFlightMode(flightMode)
                .setOrientation(orientationX, orientationY, orientationZ, orientationW);
    }

    /**
     * 1件のテスト用Commandオブジェクトのリストを生成する。
     */
    public static List<Command> newSingleCommands(Generator generator, CommandType commandType) {
        CommandId id = generator.newCommandId();
        LocalDateTime time = generator.newTime();
        return Arrays.asList(new Command[] {
            new Command(id, commandType, time)
        });
    }

    /**
     * 複数件のテスト用Commandオブジェクトのリストを生成する。
     */
    public static List<Command> newSeveralCommands(Generator generator) {
        return Arrays.asList(new Command[] {
            new Command(
                    generator.newCommandId(),
                    CommandType.ARM,
                    generator.newTime()),
            new Command(
                    generator.newCommandId(),
                    CommandType.DISARM,
                    generator.newTime()),
            new Command(
                    generator.newCommandId(),
                    CommandType.UPLOAD,
                    generator.newTime()),
        });
    }

    /**
     * 1件のテスト用UploadMissionオブジェクトのリストを生成する。
     */
    public static List<UploadMission> newSingleUploadMissions(Generator generator) {
        CommandId id = generator.newCommandId();
        return Arrays.asList(new UploadMission[] {
            new UploadMission(id, new MissionId("MISSION_ID_SAMPLE_1"))
        });
    }

    /**
     * 複数件のテスト用UploadMissionオブジェクトのリストを生成する。
     */
    public static List<UploadMission> newSeveralUploadMissions(Generator generator) {
        return Arrays.asList(new UploadMission[] {
            new UploadMission(
                    generator.newCommandId(),
                    new MissionId("MISSION_ID_SAMPLE_1")),
            new UploadMission(
                    generator.newCommandId(),
                    new MissionId("MISSION_ID_SAMPLE_2")),
            new UploadMission(
                    generator.newCommandId(),
                    new MissionId("MISSION_ID_SAMPLE_3")),
        });
    }
}