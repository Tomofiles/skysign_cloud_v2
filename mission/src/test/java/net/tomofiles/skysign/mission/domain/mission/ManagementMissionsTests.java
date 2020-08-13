package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static com.google.common.truth.Truth.assertThat;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;
import java.util.function.Supplier;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigation;

public class ManagementMissionsTests {

    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION1 = new Version(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION2 = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            private int count = 0;

            @Override
            public MissionId newMissionId() {
                return DEFAULT_MISSION_ID;
            }

            @Override
            public Version newVersion() {
                if (count == 0) {
                    count++;
                    return DEFAULT_VERSION1;
                } else {
                    return DEFAULT_VERSION2;
                }
            }
        };
    };

    @Mock
    private MissionRepository repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Userが、新しいMissionエンティティを作成する。<br>
     * Missionエンティティの初期状態を検証する。
     */
    @Test
    public void createNewMissionTest() {
        Mission mission = MissionFactory.newInstance(DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(mission.getId()).isEqualTo(DEFAULT_MISSION_ID),
            () -> assertThat(mission.getMissionName()).isNull(),
            () -> assertThat(mission.getNavigation()).isNull(),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION1)
        );
    }

    /**
     * Userが、新しいMissionエンティティに対してMission Nameを付与する。
     */
    @Test
    public void changeMissionsNameTest() {
        Mission mission = MissionFactory.newInstance(DEFAULT_GENERATOR.get());

        String newMissionName = "new mission";
        mission.nameMission(newMissionName);

        assertAll(
            () -> assertThat(mission.getMissionName()).isEqualTo(newMissionName),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION2)
        );
    }

    /**
     * Userが、Missionエンティティに新たなNavigationを追加する。<br>
     * Navigation自体の詳細なテストは、MissionNavigationTestsにまとめている。
     */
    @Test
    public void addNavigationToMissionTest() {
        Mission mission = MissionFactory.newInstance(DEFAULT_GENERATOR.get());

        mission.replaceNavigationWith(newSingleNavigation());

        assertAll(
            () -> assertThat(mission.getNavigation()).isEqualTo(newSingleNavigation()),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION2)
        );
    }
}