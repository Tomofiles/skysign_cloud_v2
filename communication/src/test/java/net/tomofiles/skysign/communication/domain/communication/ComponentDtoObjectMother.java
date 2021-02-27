package net.tomofiles.skysign.communication.domain.communication;

import java.util.Arrays;
import java.util.List;

import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.UploadMissionComponentDto;

public class ComponentDtoObjectMother {

    /**
     * テスト用CommunicationエンティティのDTOコンポーネントを生成する。
     */
    public static CommunicationComponentDto newNormalCommunicationComponentDto(
            CommunicationId communicationId,
            boolean controlled,
            Generator generatorCommand,
            Generator generatorUploadMission) {
        return new CommunicationComponentDto(
                communicationId.getId(),
                controlled,
                newNormalTelemetryComponentDto(),
                newSeveralCommandsComponentDto(generatorCommand),
                newSeveralUploadMissionsComponentDto(generatorUploadMission)
        );
    }

    /**
     * テスト用TelemetryオブジェクトのDTOコンポーネントを生成する。
     */
    public static TelemetryComponentDto newNormalTelemetryComponentDto() {
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

        return new TelemetryComponentDto(
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
                orientationW);
    }

    /**
     * 1件のテスト用CommandオブジェクトのDTOコンポーネントを生成する。
     */
    public static CommandComponentDto newSingleCommandComponentDto(Generator generator, CommandType type) {
        return new CommandComponentDto(
                generator.newCommandId().getId(),
                type.name(),
                generator.newTime());
    }

    /**
     * 複数件のテスト用CommandオブジェクトのDTOコンポーネントのリストを生成する。
     */
    public static List<CommandComponentDto> newSeveralCommandsComponentDto(Generator generator) {
        return Arrays.asList(new CommandComponentDto[] {
                new CommandComponentDto(
                        generator.newCommandId().getId(),
                        CommandType.ARM.name(),
                        generator.newTime()),
                new CommandComponentDto(
                        generator.newCommandId().getId(),
                        CommandType.DISARM.name(),
                        generator.newTime()),
                new CommandComponentDto(
                        generator.newCommandId().getId(),
                        CommandType.UPLOAD.name(),
                        generator.newTime()),
        });
    }

    /**
     * 1件のテスト用UploadMissionオブジェクトのDTOコンポーネントを生成する。
     */
    public static UploadMissionComponentDto newSingleUploadMissionComponentDto(Generator generator, MissionId missionId) {
        return new UploadMissionComponentDto(
                generator.newCommandId().getId(),
                missionId.getId());
    }

    /**
     * 複数件のテスト用UploadMissionオブジェクトのDTOコンポーネントのリストを生成する。
     */
    public static List<UploadMissionComponentDto> newSeveralUploadMissionsComponentDto(Generator generator) {
        return Arrays.asList(new UploadMissionComponentDto[] {
                new UploadMissionComponentDto(
                        generator.newCommandId().getId(),
                        "MISSION_ID_SAMPLE_1"),
                new UploadMissionComponentDto(
                        generator.newCommandId().getId(),
                        "MISSION_ID_SAMPLE_2"),
                new UploadMissionComponentDto(
                        generator.newCommandId().getId(),
                        "MISSION_ID_SAMPLE_3"),
        });
    }
}