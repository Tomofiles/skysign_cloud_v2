package net.tomofiles.skysign.communication.infra.communication;

import java.util.Arrays;
import java.util.List;

import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.Generator;

public class RecordObjectMother {
    
    /**
     * 通常のCommunicationレコードを生成する。
     */
    public static CommunicationRecord newNormalCommunicationRecord(CommunicationId id) {
        return new CommunicationRecord(id.getId());
    }
    
    /**
     * 通常のTelemetryレコードを生成する。
     */
    public static TelemetryRecord newNormalTelemetryRecord(CommunicationId id) {
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

        return new TelemetryRecord(
                id.getId(),
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
                orientationW
        );
    }

    /**
     * 空のTelemetryレコードを生成する。
     */
    public static TelemetryRecord newEmptyTelemetryRecord(CommunicationId id) {
        TelemetryRecord telemetry = new TelemetryRecord();
        telemetry.setCommId(id.getId());
        return telemetry;
    }

    /**
     * 1件のCommandレコードを生成する。
     */
    public static CommandRecord newSingleCommandRecord(CommunicationId id, Generator generator) {
        return new CommandRecord(
                generator.newCommandId().getId(),
                id.getId(),
                CommandType.ARM.name(),
                generator.newTime());
    }

    /**
     * 複数件のCommandレコードを生成する。
     */
    public static List<CommandRecord> newSeveralCommandRecords(CommunicationId id, Generator generator) {
        return Arrays.asList(new CommandRecord[] {
                new CommandRecord(
                        generator.newCommandId().getId(),
                        id.getId(),
                        CommandType.ARM.name(),
                        generator.newTime()),
                new CommandRecord(
                        generator.newCommandId().getId(),
                        id.getId(),
                        CommandType.DISARM.name(),
                        generator.newTime()),
                new CommandRecord(
                        generator.newCommandId().getId(),
                        id.getId(),
                        CommandType.UPLOAD.name(),
                        generator.newTime()),
        });
    }

    /**
     * 1件のUploadMissionレコードを生成する。
     */
    public static UploadMissionRecord newSingleUploadMissionRecord(CommunicationId id, Generator generator) {
        return new UploadMissionRecord(
                generator.newCommandId().getId(),
                id.getId(),
                "MISSION_ID_SAMPLE_1");
    }

    /**
     * 複数件のUploadMissionレコードを生成する。
     */
    public static List<UploadMissionRecord> newSeveralUploadMissionRecords(CommunicationId id, Generator generator) {
        return Arrays.asList(new UploadMissionRecord[] {
                new UploadMissionRecord(
                        generator.newCommandId().getId(),
                        id.getId(),
                        "MISSION_ID_SAMPLE_1"),
                new UploadMissionRecord(
                        generator.newCommandId().getId(),
                        id.getId(),
                        "MISSION_ID_SAMPLE_2"),
                new UploadMissionRecord(
                        generator.newCommandId().getId(),
                        id.getId(),
                        "MISSION_ID_SAMPLE_3"),
        });
    }
}