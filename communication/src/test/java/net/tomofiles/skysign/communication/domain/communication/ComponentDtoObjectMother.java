package net.tomofiles.skysign.communication.domain.communication;

import java.util.Arrays;
import java.util.List;

import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;

public class ComponentDtoObjectMother {

    /**
     * テスト用CommunicationエンティティのDTOコンポーネントを生成する。
     */
    public static CommunicationComponentDto newNormalCommunicationComponentDto(CommunicationId communicationId, MissionId missionId, Generator generator) {
        return new CommunicationComponentDto(
                communicationId.getId(),
                missionId.getId(),
                newNormalTelemetryComponentDto(),
                newSeveralCommandsComponentDto(generator)
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
}