package net.tomofiles.skysign.mission.domain.mission;

import java.util.Arrays;

import net.tomofiles.skysign.mission.domain.mission.component.MissionComponentDto;
import net.tomofiles.skysign.mission.domain.mission.component.WaypointComponentDto;

public class ComponentDtoObjectMother {

    /**
     * 昇順のWaypointを複数件持つNavigationを含むMissionエンティティのDTOコンポーネントを生成する。
     */
    public static MissionComponentDto newSeveralNavigationMissionComponentDto(Generator generator) {
        Version version = generator.newVersion();

        return new MissionComponentDto(
                generator.newMissionId().getId(),
                "mission name",
                0.0,
                version.getVersion(),
                version.getVersion(),
                Arrays.asList(new WaypointComponentDto[] {
                        new WaypointComponentDto(
                                1,
                                1.0,
                                2.0,
                                3.0,
                                4.0),
                        new WaypointComponentDto(
                                2,
                                11.0,
                                12.0,
                                13.0,
                                14.0),
                        new WaypointComponentDto(
                                3,
                                21.0,
                                22.0,
                                23.0,
                                24.0),
                })
        );
    }

}