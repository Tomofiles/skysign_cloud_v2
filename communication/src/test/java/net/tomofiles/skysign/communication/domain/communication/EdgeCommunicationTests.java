package net.tomofiles.skysign.communication.domain.communication;

import static com.google.common.truth.Truth.assertThat;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.function.Supplier;

import static net.tomofiles.skysign.communication.domain.communication.SnapshotObjectMother.newNormalTelemetrySnapshot;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newSingleCommandCommunication;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newSeveralCommandsCommunication;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalTelemetry;

public class EdgeCommunicationTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID = new CommandId(UUID.randomUUID().toString());
    private static final MissionId DEFAULT_MISSION_ID = new MissionId("MISSION_ID_SAMPLE_1");
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
    private static final CommandId DEFAULT_COMMAND_ID1 = new CommandId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID2 = new CommandId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID3 = new CommandId(UUID.randomUUID().toString());
    private static final LocalDateTime DEFAULT_COMMAND_TIME1 = LocalDateTime.of(2020, 07, 22, 10, 30, 25);
    private static final LocalDateTime DEFAULT_COMMAND_TIME2 = LocalDateTime.of(2020, 07, 22, 10, 30, 30);
    private static final LocalDateTime DEFAULT_COMMAND_TIME3 = LocalDateTime.of(2020, 07, 22, 10, 30, 45);
    private static final Supplier<Generator> DEFAULT_GENERATOR_IN_RONDOM_ORDER = () -> {
        return new Generator(){
            private List<CommandId> commandIds = new ArrayList<>(Arrays.asList(new CommandId[] {
                    DEFAULT_COMMAND_ID1,
                    DEFAULT_COMMAND_ID2,
                    DEFAULT_COMMAND_ID3
            }));
            private List<LocalDateTime> times = new ArrayList<>(Arrays.asList(new LocalDateTime[] {
                    DEFAULT_COMMAND_TIME3, // 順不同
                    DEFAULT_COMMAND_TIME1, // 順不同
                    DEFAULT_COMMAND_TIME2  // 順不同
            }));
            @Override
            public CommandId newCommandId() {
                return commandIds.remove(0);
            }
            @Override
            public LocalDateTime newTime() {
                return times.remove(0);
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
     * Edgeが、既存のCommunicationエンティティのTelemetryを更新する。<br>
     * すべてのTelemetryのフィールドが更新されることを検証する。
     */
    @Test
    public void pushTelemetryToCommunicationTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);

        communication.pushTelemetry(newNormalTelemetrySnapshot());

        assertThat(communication.getTelemetry()).isEqualTo(newNormalTelemetry());
    }

    /**
     * Edgeが、既存のCommunicationエンティティからCommandIDリストを取得する。<br>
     * CommandIDはCommandをEdgeから古い順でCloudに取得しに来るため、<br>
     * CommandのTimeの昇順でソートされていることを検証する。
     */
    @Test
    public void pullCommandIDsFromCommunicationTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSeveralCommandsCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR_IN_RONDOM_ORDER.get(),
                        DEFAULT_GENERATOR_IN_RONDOM_ORDER.get(),
                        DEFAULT_GENERATOR_IN_RONDOM_ORDER.get()));

        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);

        List<CommandId> commandIds = communication.getCommandIds();

        assertThat(commandIds).isEqualTo(Arrays.asList(new CommandId[] {
            DEFAULT_COMMAND_ID2,
            DEFAULT_COMMAND_ID3,
            DEFAULT_COMMAND_ID1
        }));
    }

    /**
     * Edgeが、既存のCommunicationエンティティからCommandを取得する。<br>
     * CommandIDに合致するCommandが返却され、Communicationエンティティから<br>
     * Commandが削除されることを検証する。
     */
    @Test
    public void pullCommandFromCommunicationTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSingleCommandCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get()));

        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);

        CommandType type = communication.pullCommandById(DEFAULT_COMMAND_ID);

        assertAll(
            () -> assertThat(type).isEqualTo(CommandType.ARM),
            () -> assertThat(communication.getCommandIds()).hasSize(0)
        );
    }

    /**
     * Edgeが、既存のCommunicationエンティティからCommandを取得する。<br>
     * Commandはまだ発行されていないので、NULLが返却されることを検証する。
     */
    @Test
    public void pullNoCommandFromCommunicationTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);

        CommandType type = communication.pullCommandById(DEFAULT_COMMAND_ID);

        assertAll(
            () -> assertThat(type).isNull(),
            () -> assertThat(communication.getCommandIds()).hasSize(0),
            () -> assertThat(communication.getCommands()).hasSize(0)
        );
    }

    /**
     * Edgeが、既存のCommunicationエンティティからUploadMissionを取得する。<br>
     * CommandIDに合致するUploadMissionが返却され、Communicationエンティティから<br>
     * UploadMissionが削除されることを検証する。
     */
    @Test
    public void pullUploadMissionFromCommunicationTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSingleCommandCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get()));

        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);

        MissionId missionId = communication.pullUploadMissionById(DEFAULT_COMMAND_ID);

        assertAll(
            () -> assertThat(missionId).isEqualTo(DEFAULT_MISSION_ID),
            () -> assertThat(communication.getUploadMissions()).hasSize(0)
        );
    }

}