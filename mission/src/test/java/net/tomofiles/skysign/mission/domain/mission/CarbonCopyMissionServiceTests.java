package net.tomofiles.skysign.mission.domain.mission;

import java.util.UUID;
import java.util.function.Supplier;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationMission;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationCarbonCopiedMission;

public class CarbonCopyMissionServiceTests {
    
    private static final MissionId ORIGINAL_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final MissionId NEW_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public MissionId newMissionId() {
                throw new IllegalStateException();
            }
            @Override
            public Version newVersion() {
                throw new IllegalStateException();
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
     * Missionのカーボンコピーを作成するサービスをテストする。
     * 指定されたIDのMissionを、指定されたIDでコピーする。
     */
    @Test
    public void carbonCopyMissionTest() {
        when(this.repository.getById(ORIGINAL_MISSION_ID))
                .thenReturn(newSingleNavigationMission(
                        ORIGINAL_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        CarbonCopyMissionService.copy(
                DEFAULT_GENERATOR.get(), 
                this.repository, 
                ORIGINAL_MISSION_ID, 
                NEW_MISSION_ID);

        ArgumentCaptor<Mission> missionCaptor = ArgumentCaptor.forClass(Mission.class);
        verify(this.repository, times(1)).save(missionCaptor.capture());

        Mission expectMission = newSingleNavigationMission(
                NEW_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get());

        assertThat(missionCaptor.getValue()).isEqualTo(expectMission);
    }

    /**
     * Missionのカーボンコピーを作成するサービスをテストする。
     * コピー後のIDのMissionがすでに存在する場合、コピーを行わず
     * 正常終了することを検証する。
     */
    @Test
    public void copySuccessWhenAlreadyExistsMissionWhenCarbonCopyMissionTest() {
        when(this.repository.getById(NEW_MISSION_ID))
                .thenReturn(newSingleNavigationCarbonCopiedMission(
                        NEW_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        CarbonCopyMissionService.copy(
                DEFAULT_GENERATOR.get(), 
                this.repository, 
                ORIGINAL_MISSION_ID, 
                NEW_MISSION_ID);

        verify(this.repository, times(0)).save(any());
    }

    /**
     * Missionのカーボンコピーを作成するサービスをテストする。
     * 指定されたIDのMissionの取得がエラーとなった場合、
     * 正常終了することを検証する。
     */
    @Test
    public void getErrorWhenCarbonCopyMissionTest() {
        CarbonCopyMissionService.copy(
                DEFAULT_GENERATOR.get(), 
                this.repository, 
                ORIGINAL_MISSION_ID, 
                NEW_MISSION_ID);

        verify(this.repository, times(0)).save(any());
    }

}