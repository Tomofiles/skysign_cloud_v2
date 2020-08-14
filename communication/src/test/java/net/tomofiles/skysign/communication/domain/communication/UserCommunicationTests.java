package net.tomofiles.skysign.communication.domain.communication;

import static com.google.common.truth.Truth.assertThat;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
import java.util.UUID;
import java.util.function.Supplier;

import static net.tomofiles.skysign.communication.domain.communication.SnapshotObjectMother.newNormalTelemetrySnapshot;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;

public class UserCommunicationTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID = new CommandId(UUID.randomUUID().toString());
    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final LocalDateTime DEFAULT_COMMAND_TIME = LocalDateTime.of(2020, 1, 1, 0, 0, 0);
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public CommandId newCommandId() {
                return DEFAULT_COMMAND_ID;
            }
            @Override
            public LocalDateTime newTime() {
                return DEFAULT_COMMAND_TIME;
            }
        };
    };

    @Mock
    private CommunicationRepository repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Userが、新しいCommunicationエンティティを作成する。<br>
     * Communicationエンティティの初期状態を検証する。
     */
    @Test
    public void createNewCommunicationTest() {
        Communication communication = CommunicationFactory.newInstance(DEFAULT_COMMUNICATION_ID, DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(communication.getId()).isEqualTo(DEFAULT_COMMUNICATION_ID),
            () -> assertThat(communication.getMissionId()).isNull(),
            () -> assertThat(communication.getCommands()).hasSize(0),
            () -> assertThat(communication.getTelemetry()).isEqualTo(Telemetry.newInstance())
        );
    }

    /**
     * Userが、既存のCommunicationエンティティにCommandを追加する。<br>
     * Commandが追加され、IDとTimeが付与されていることを検証する。
     */
    @Test
    public void pushCommandToCommunicationTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(DEFAULT_COMMUNICATION_ID, DEFAULT_GENERATOR.get()));

        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);

        communication.pushCommand(CommandType.ARM);

        assertAll(
            () -> assertThat(communication.getCommandIds()).hasSize(1),
            () -> assertThat(communication.getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID),
            () -> assertThat(communication.getCommands()).hasSize(1),
            () -> assertThat(communication.getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID),
            () -> assertThat(communication.getCommands().get(0).getType()).isEqualTo(CommandType.ARM),
            () -> assertThat(communication.getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME)
        );
    }

    /**
     * Userが、既存のCommunicationエンティティからTelemetryを取得する。<br>
     * Telemetryのスナップショットが生成され、返却されることを検証する。
     */
    @Test
    public void pullTelemetryFromCommunicationTest() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        Communication communication = repository.getById(DEFAULT_COMMUNICATION_ID);

        TelemetrySnapshot telemetry = communication.pullTelemetry();

        assertThat(telemetry).isEqualTo(newNormalTelemetrySnapshot());
    }
}