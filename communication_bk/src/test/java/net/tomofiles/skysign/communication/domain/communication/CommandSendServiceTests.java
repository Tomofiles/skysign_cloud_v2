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

public class CommandSendServiceTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID1 = new CommandId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID2 = new CommandId(UUID.randomUUID().toString());
    private static final LocalDateTime DEFAULT_COMMAND_TIME1 = LocalDateTime.of(2020, 1, 1, 0, 0, 0);
    private static final LocalDateTime DEFAULT_COMMAND_TIME2 = LocalDateTime.of(2020, 1, 1, 0, 0, 1);
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            int commandCount, timeCount = 0;
            @Override
            public CommandId newCommandId() {
                switch (commandCount) {
                    case 0:
                        commandCount++;
                        return DEFAULT_COMMAND_ID1;
                    case 1:
                        commandCount++;
                        return DEFAULT_COMMAND_ID2;
                    default:
                        throw new RuntimeException();
                }
            }
            @Override
            public LocalDateTime newTime() {
                switch (timeCount) {
                    case 0:
                        timeCount++;
                        return DEFAULT_COMMAND_TIME1;
                    case 1:
                        timeCount++;
                        return DEFAULT_COMMAND_TIME2;
                    default:
                        throw new RuntimeException();
                }
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
    public void notFoundCommunicationWhenSendCommandTest() {
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.ARM);

        verify(this.repository, times(0)).save(any());
        verify(this.serviceReturn, times(0)).setCommunication(any());
        verify(this.serviceReturn, times(0)).setCommandId(any());
    }

    /**
     * ARMコマンドを送信する。
     */
    @Test
    public void sendArmCommandTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.ARM);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.ARM),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * DISARMコマンドを送信する。
     */
    @Test
    public void sendDisarmCommandTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.DISARM);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.DISARM),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * STARTコマンドを送信する。<br>
     * TelemetryがArmed状態の場合、STARTコマンドのみ送信されることを検証する。
     */
    @Test
    public void sendStartCommandTest() {
        Communication communication = CommunicationFactory.newInstance(
            DEFAULT_COMMUNICATION_ID,
            DEFAULT_GENERATOR.get());
        communication.pushTelemetry(new TelemetrySnapshot(0.0, 0.0, 0.0, 0.0, 0.0, true, "", 0.0, 0.0, 0.0, 0.0));
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(communication);
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.START);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.START),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * STARTコマンドを送信する。<br>
     * TelemetryがDisarmed状態の場合、ArmコマンドとStartコマンドの両方が
     * 送信されることを検証する。
     */
    @Test
    public void sendArmAndStartCommandTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.START);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(2),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(1)).isEqualTo(DEFAULT_COMMAND_ID2),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(2),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.ARM),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(1).getId()).isEqualTo(DEFAULT_COMMAND_ID2),
            () -> assertThat(saveCaptor.getValue().getCommands().get(1).getType()).isEqualTo(CommandType.START),
            () -> assertThat(saveCaptor.getValue().getCommands().get(1).getTime()).isEqualTo(DEFAULT_COMMAND_TIME2),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(1)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * PAUSEコマンドを送信する。
     */
    @Test
    public void sendPauseCommandTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.PAUSE);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.PAUSE),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * TAKEOFFコマンドを送信する。<br>
     * TelemetryがArmed状態の場合、TAKEOFFコマンドのみ送信されることを検証する。
     */
    @Test
    public void sendTakeoffCommandTest() {
        Communication communication = CommunicationFactory.newInstance(
            DEFAULT_COMMUNICATION_ID,
            DEFAULT_GENERATOR.get());
        communication.pushTelemetry(new TelemetrySnapshot(0.0, 0.0, 0.0, 0.0, 0.0, true, "", 0.0, 0.0, 0.0, 0.0));
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(communication);
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.TAKEOFF);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.TAKEOFF),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * TAKEOFFコマンドを送信する。<br>
     * TelemetryがDisarmed状態の場合、ArmコマンドとTAKEOFFコマンドの両方が
     * 送信されることを検証する。
     */
    @Test
    public void sendArmAndTakeoffCommandTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.TAKEOFF);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(2),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(1)).isEqualTo(DEFAULT_COMMAND_ID2),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(2),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.ARM),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(1).getId()).isEqualTo(DEFAULT_COMMAND_ID2),
            () -> assertThat(saveCaptor.getValue().getCommands().get(1).getType()).isEqualTo(CommandType.TAKEOFF),
            () -> assertThat(saveCaptor.getValue().getCommands().get(1).getTime()).isEqualTo(DEFAULT_COMMAND_TIME2),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(1)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * LANDコマンドを送信する。
     */
    @Test
    public void sendLandCommandTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.LAND);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.LAND),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }

    /**
     * RETURNコマンドを送信する。
     */
    @Test
    public void sendReturnCommandTest() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));
        
        CommandSendService.send(
            this.repository, 
            serviceReturn::setCommunication,
            serviceReturn::setCommandId,
            DEFAULT_COMMUNICATION_ID, 
            CommandType.RETURN);

        ArgumentCaptor<Communication> saveCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(saveCaptor.capture());

        ArgumentCaptor<Communication> setCommunicationCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.serviceReturn, times(1)).setCommunication(setCommunicationCaptor.capture());
        ArgumentCaptor<CommandId> setCommandIdCaptor = ArgumentCaptor.forClass(CommandId.class);
        verify(this.serviceReturn, times(1)).setCommandId(setCommandIdCaptor.capture());

        assertAll(
            () -> assertThat(saveCaptor.getValue().getCommandIds()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands()).hasSize(1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getId()).isEqualTo(DEFAULT_COMMAND_ID1),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getType()).isEqualTo(CommandType.RETURN),
            () -> assertThat(saveCaptor.getValue().getCommands().get(0).getTime()).isEqualTo(DEFAULT_COMMAND_TIME1),
            () -> assertThat(saveCaptor.getValue()).isEqualTo(setCommunicationCaptor.getValue()),
            () -> assertThat(saveCaptor.getValue().getCommandIds().get(0)).isEqualTo(setCommandIdCaptor.getValue())
        );
    }
}