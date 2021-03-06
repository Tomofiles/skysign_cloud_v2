package net.tomofiles.skysign.mission.api;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;
import java.util.function.Supplier;

import net.tomofiles.skysign.mission.domain.mission.Generator;
import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.domain.mission.Version;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.MissionRepository;
import net.tomofiles.skysign.mission.service.ManageMissionService;

import static net.tomofiles.skysign.mission.api.EventObjectMother.newNormalMissionCopiedWhenCopiedEvent;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationMission;

public class MissionEventHandlerTests {
    
    private static final MissionId ORIGINAL_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final MissionId NEW_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final String EXCHANGE_NAME_COPIED_EVENT = "exchange_name_copied_event";
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

    @InjectMocks
    private ManageMissionService service;

    private MissionEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new MissionEventHandler(this.service);
        this.eventHandler.setEXCHANGE_NAME_COPIED_EVENT(EXCHANGE_NAME_COPIED_EVENT);
    }

    /**
     * Flightplanがコピーされたときに新たなMissionIDが発行されたイベントを
     * 受信した場合の処理を確認する。<br>
     * 対象のMissionのカーボンコピーが作成されたことを検証する。
     */
    @Test
    public void fireMissionCopiedWhenFlightplanCopiedEvent() throws Exception {
        when(this.repository.getById(ORIGINAL_MISSION_ID))
                .thenReturn(newSingleNavigationMission(
                        ORIGINAL_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        this.eventHandler.processMissionCopiedWhenFlightplanCopiedEvent(
            newNormalMissionCopiedWhenCopiedEvent(
                ORIGINAL_MISSION_ID,
                NEW_MISSION_ID
            ));

        ArgumentCaptor<Mission> missionCaptor = ArgumentCaptor.forClass(Mission.class);
        verify(this.repository, times(1)).save(missionCaptor.capture());

        assertThat(missionCaptor.getValue().getId()).isEqualTo(NEW_MISSION_ID);
    }
}