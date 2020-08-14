package net.tomofiles.skysign.communication.domain.communication;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;

public class CommunicationObjectMother {
    
    /**
     * テスト用Communicationエンティティを生成する。
     */
    public static Communication newNormalCommunication(CommunicationId communicationId, MissionId missionId, Generator generator) {
        Communication communication = CommunicationFactory.newInstance(communicationId, generator);
        communication.setMissionId(missionId);
        communication.setTelemetry(newNormalTelemetry());
        return communication;
    }

    /**
     * 1件のCommandを持つテスト用Communicationエンティティを生成する。
     */
    public static Communication newSingleCommandCommunication(CommunicationId communicationId, MissionId missionId, Generator generator) {
        Communication communication = CommunicationFactory.newInstance(communicationId, generator);
        communication.setMissionId(missionId);
        communication.getCommands().addAll(newSingleCommands(generator, CommandType.ARM));
        communication.setTelemetry(newNormalTelemetry());
        return communication;
    }

    /**
     * 複数件のCommandリストを持つテスト用Communicationエンティティを生成する。
     */
    public static Communication newSeveralCommandsCommunication(CommunicationId communicationId, MissionId missionId, Generator generator) {
        Communication communication = CommunicationFactory.newInstance(communicationId, generator);
        communication.setMissionId(missionId);
        communication.getCommands().addAll(newSeveralCommands(generator));
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
}