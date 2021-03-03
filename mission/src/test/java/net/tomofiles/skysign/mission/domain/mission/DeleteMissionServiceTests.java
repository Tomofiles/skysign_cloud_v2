package net.tomofiles.skysign.mission.domain.mission;

import java.util.UUID;
import java.util.function.Supplier;

import static com.google.common.truth.Truth.assertThat;
import static org.junit.Assert.assertThrows;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.MockitoAnnotations.initMocks;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationMission;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationCarbonCopiedMission;

public class DeleteMissionServiceTests {
    
    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public MissionId newMissionId() {
                return DEFAULT_MISSION_ID;
            }

            @Override
            public Version newVersion() {
                return DEFAULT_VERSION;
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
     * Missionを削除するサービスをテストする。
     */
    @Test
    public void deleteMissionServiceTest() {
        Mission mission = newSingleNavigationMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get());

        DeleteMissionService.delete(repository, mission);

        verify(repository, times(1)).remove(DEFAULT_MISSION_ID, DEFAULT_VERSION);
    }

    /**
     * Missionを削除するサービスをテストする。
     * 対象のMissionが存在しない場合、エラーが発生せず
     * 正常終了することを検証する。
     */
    @Test
    public void noMissionWhenDeleteMissionServiceTest() {
        DeleteMissionService.delete(repository, null);

        verify(repository, times(0)).remove(any(), any());
    }

    /**
     * Missionを削除するサービスをテストする。
     * カーボンコピーされたMissionの削除時に例外がスローされることを検証する。
     */
    @Test
    public void deleteCarbonCopiedMissionWhenDeleteMissionServiceTest() {
        Mission mission = newSingleNavigationCarbonCopiedMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get());

        CannotChangeMissionException e
                = assertThrows(
                        CannotChangeMissionException.class,
                        () -> DeleteMissionService.delete(repository, mission));
        
        assertThat(e).hasMessageThat().contains("cannot delete carbon copied mission");
    }

}