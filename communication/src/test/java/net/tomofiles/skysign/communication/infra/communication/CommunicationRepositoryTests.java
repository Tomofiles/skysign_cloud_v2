package net.tomofiles.skysign.communication.infra.communication;

import static com.google.common.truth.Truth.assertThat;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.function.Supplier;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.Generator;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;

import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newSingleCommandCommunication;
import static net.tomofiles.skysign.communication.domain.communication.ComponentDtoObjectMother.newNormalCommunicationComponentDto;
import static net.tomofiles.skysign.communication.infra.communication.RecordObjectMother.newNormalCommunicationRecord;
import static net.tomofiles.skysign.communication.infra.communication.RecordObjectMother.newNormalTelemetryRecord;
import static net.tomofiles.skysign.communication.infra.communication.RecordObjectMother.newEmptyTelemetryRecord;
import static net.tomofiles.skysign.communication.infra.communication.RecordObjectMother.newSingleCommandRecord;
import static net.tomofiles.skysign.communication.infra.communication.RecordObjectMother.newSeveralCommandRecords;
import static net.tomofiles.skysign.communication.infra.communication.RecordObjectMother.newSingleUploadMissionRecord;
import static net.tomofiles.skysign.communication.infra.communication.RecordObjectMother.newSeveralUploadMissionRecords;

public class CommunicationRepositoryTests {

    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final boolean DEFAULT_CONTROLLED = true;
    private static final CommandId DEFAULT_COMMAND_ID1 = new CommandId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID2 = new CommandId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID3 = new CommandId(UUID.randomUUID().toString());
    private static final LocalDateTime DEFAULT_COMMAND_TIME1 = LocalDateTime.of(2020, 07, 22, 10, 30, 25);
    private static final LocalDateTime DEFAULT_COMMAND_TIME2 = LocalDateTime.of(2020, 07, 22, 10, 30, 30);
    private static final LocalDateTime DEFAULT_COMMAND_TIME3 = LocalDateTime.of(2020, 07, 22, 10, 30, 45);
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            private List<CommandId> commandIds = new ArrayList<>(Arrays.asList(new CommandId[] {
                    DEFAULT_COMMAND_ID1,
                    DEFAULT_COMMAND_ID2,
                    DEFAULT_COMMAND_ID3
            }));
            private List<LocalDateTime> times = new ArrayList<>(Arrays.asList(new LocalDateTime[] {
                    DEFAULT_COMMAND_TIME1,
                    DEFAULT_COMMAND_TIME2,
                    DEFAULT_COMMAND_TIME3
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
    private static final Supplier<Generator> DEFAULT_GENERATOR_SINGLE_1 = () -> {
        return new Generator(){
            @Override
            public CommandId newCommandId() {
                return DEFAULT_COMMAND_ID1;
            }
            @Override
            public LocalDateTime newTime() {
                return DEFAULT_COMMAND_TIME1;
            }
        };
    };
    private static final Supplier<Generator> DEFAULT_GENERATOR_SINGLE_2 = () -> {
        return new Generator(){
            @Override
            public CommandId newCommandId() {
                return DEFAULT_COMMAND_ID2;
            }
            @Override
            public LocalDateTime newTime() {
                return DEFAULT_COMMAND_TIME2;
            }
        };
    };

    @Mock
    private CommunicationMapper communicationMapper;

    @Mock
    private TelemetryMapper telemetryMapper;

    @Mock
    private CommandMapper commandMapper;

    @Mock
    private UploadMissionMapper uploadMissionMapper;

    @InjectMocks
    private CommunicationRepositoryImpl repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * リポジトリーからCommunicationエンティティをすべて取得する。
     */
    @Test
    public void getAllCommunicationsTest() {
        when(this.communicationMapper.findAll())
                .thenReturn(Arrays.asList(new CommunicationRecord[] {
                        newNormalCommunicationRecord(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_CONTROLLED),
                        newNormalCommunicationRecord(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_CONTROLLED),
                        newNormalCommunicationRecord(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_CONTROLLED)
                }));
        when(this.telemetryMapper.find(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newNormalTelemetryRecord(DEFAULT_COMMUNICATION_ID));
        when(this.commandMapper.findByCommId(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newSeveralCommandRecords(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        when(this.uploadMissionMapper.findByCommId(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newSeveralUploadMissionRecords(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        List<Communication> communications = this.repository.getAll();

        CommunicationComponentDto dto = CommunicationFactory.takeApart(communications.get(0));
        CommunicationComponentDto expectDto = newNormalCommunicationComponentDto(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_CONTROLLED,
                DEFAULT_GENERATOR.get(),
                DEFAULT_GENERATOR.get());
        
        assertAll(
            () -> assertThat(communications).hasSize(3),
            () -> assertThat(dto).isEqualTo(expectDto)
        );
    }

    /**
     * リポジトリーからCommunicationエンティティを一つ取得する。<br>
     * 対象のエンティティが存在しない場合、NULLが返却されることを検証する。
     */
    @Test
    public void getAllNoCommunicationTest() {
        List<Communication> communications = this.repository.getAll();

        assertThat(communications).hasSize(0);
    }

    /**
     * リポジトリーからCommunicationエンティティを一つ取得する。
     */
    @Test
    public void getCommunicationByIdTest() {
        when(this.communicationMapper.find(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newNormalCommunicationRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_CONTROLLED));
        when(this.telemetryMapper.find(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newNormalTelemetryRecord(
                        DEFAULT_COMMUNICATION_ID));
        when(this.commandMapper.findByCommId(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newSeveralCommandRecords(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        when(this.uploadMissionMapper.findByCommId(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newSeveralUploadMissionRecords(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);
        
        CommunicationComponentDto dto = CommunicationFactory.takeApart(communication);
        CommunicationComponentDto expectDto = newNormalCommunicationComponentDto(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_CONTROLLED,
                DEFAULT_GENERATOR.get(),
                DEFAULT_GENERATOR.get());

        assertThat(dto).isEqualTo(expectDto);
    }

    /**
     * リポジトリーからCommunicationエンティティを一つ取得する。<br>
     * 対象のエンティティが存在しない場合、NULLが返却されることを検証する。
     */
    @Test
    public void getNoCommunicationByIdTest() {
        Communication communication = this.repository.getById(DEFAULT_COMMUNICATION_ID);

        assertThat(communication).isNull();
    }

    /**
     * リポジトリーにCommunicationエンティティを一つ保存する。<br>
     * 既存のエンティティが無いため、新規登録されることを検証する。
     */
    @Test
    public void saveNewCommunicationTest() {
        this.repository.save(newSingleCommandCommunication(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_CONTROLLED,
                DEFAULT_GENERATOR.get(),
                DEFAULT_GENERATOR_SINGLE_1.get(),
                DEFAULT_GENERATOR_SINGLE_1.get()));

        verify(this.communicationMapper, times(1))
                .create(newNormalCommunicationRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_CONTROLLED));
        verify(this.telemetryMapper, times(1))
                .create(newNormalTelemetryRecord(DEFAULT_COMMUNICATION_ID));
        verify(this.commandMapper, times(1))
                .create(newSingleCommandRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR_SINGLE_1.get()));
        verify(this.uploadMissionMapper, times(1))
                .create(newSingleUploadMissionRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
    }

    /**
     * リポジトリーにCommunicationエンティティを一つ保存する。<br>
     * 既存のエンティティが存在するため、更新されることを検証する。
     */
    @Test
    public void savePreExistCommunicationTest() {
        when(this.communicationMapper.find(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newNormalCommunicationRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_CONTROLLED));
        when(this.telemetryMapper.find(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(newEmptyTelemetryRecord(DEFAULT_COMMUNICATION_ID));
        when(this.commandMapper.findByCommId(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(Arrays.asList(new CommandRecord[] {
                        newSingleCommandRecord(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_GENERATOR_SINGLE_1.get())
                }));
        when(this.uploadMissionMapper.findByCommId(DEFAULT_COMMUNICATION_ID.getId()))
                .thenReturn(Arrays.asList(new UploadMissionRecord[] {
                        newSingleUploadMissionRecord(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_GENERATOR_SINGLE_1.get())
                }));

        this.repository.save(newSingleCommandCommunication(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_CONTROLLED,
                DEFAULT_GENERATOR.get(),
                DEFAULT_GENERATOR_SINGLE_2.get(),
                DEFAULT_GENERATOR_SINGLE_2.get()));

        verify(this.communicationMapper, times(1))
                .update(newNormalCommunicationRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_CONTROLLED));
        verify(this.telemetryMapper, times(1))
                .update(newNormalTelemetryRecord(DEFAULT_COMMUNICATION_ID));
        verify(this.commandMapper, times(1))
                .create(newSingleCommandRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR_SINGLE_2.get()));
        verify(this.commandMapper, times(1))
                .delete(DEFAULT_COMMAND_ID1.getId());
        verify(this.uploadMissionMapper, times(1))
                .create(newSingleUploadMissionRecord(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR_SINGLE_2.get()));
        verify(this.uploadMissionMapper, times(1))
                .delete(DEFAULT_COMMAND_ID1.getId());
    }

    /**
     * リポジトリーからCommunicationエンティティを一つ削除する。
     */
    @Test
    public void removeCommunicationTest() {
        this.repository.remove(DEFAULT_COMMUNICATION_ID);

        verify(this.communicationMapper, times(1)).delete(DEFAULT_COMMUNICATION_ID.getId());
    }
}