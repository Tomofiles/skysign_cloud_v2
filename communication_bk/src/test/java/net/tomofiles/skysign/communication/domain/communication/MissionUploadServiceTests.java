package net.tomofiles.skysign.communication.domain.communication;

import static com.google.common.truth.Truth.assertThat;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
import java.util.UUID;
import java.util.function.Supplier;

public class MissionUploadServiceTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID = new CommandId(UUID.randomUUID().toString());
    private static final LocalDateTime DEFAULT_COMMAND_TIME = LocalDateTime.of(2020, 1, 1, 0, 0, 0);
    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
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

    private static interface CommandSendServiceReturn {
        public void setCommunication(Communication communication);
        public void setCommandId(CommandId commandId);
    }

    @Mock
    private CommunicationRepository repository;

    @Mock
    private CommandSendServiceReturn serviceReturn;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Communicationが存在しない場合、エラーが発生せず正常終了する。
     */
    @Test
    public void notFoundCommunicationWhenUploadMissionTest() {
        MissionUploadService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            DEFAULT_MISSION_ID);

        verify(this.repository, times(0)).save(any());
        verify(this.serviceReturn, times(0)).setCommunication(any());
        verify(this.serviceReturn, times(0)).setCommandId(any());
    }

    /**
     * MissionUploadコマンドを送信する。
     */
    @Test
    public void uploadMissionTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        MissionUploadService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            DEFAULT_MISSION_ID);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.UPLOAD),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME),
            () -> assertThat(saveCaptor.getValue().getUploadMissions()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getUploadMissions().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID),
            () -> assertThat(saveCaptor.getValue().getUploadMissions().get(0).getMissionId()).isEqualTo(DEFAULT_MISSION_ID),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }
}