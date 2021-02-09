package net.tomofiles.skysign.communication.domain.communication;

import static com.google.common.truth.Truth.assertThat;
import org.junit.jupiter.api.Test;

import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.function.Supplier;

import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newSeveralCommandsCommunication;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newSeveralCommands;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalTelemetry;
import static net.tomofiles.skysign.communication.domain.communication.ComponentDtoObjectMother.newNormalCommunicationComponentDto;
import static org.junit.jupiter.api.Assertions.assertAll;

public class ComponentCommunicationTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
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
    
    /**
     * DTOからCommunicationエンティティを組み立てる。
     */
    @Test
    public void assembleIntoCommunicationTest() {
        Communication communication = CommunicationFactory.assembleFrom(
                newNormalCommunicationComponentDto(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_GENERATOR.get()
                ),
                DEFAULT_GENERATOR.get()
        );

        assertAll(
            () -> assertThat(communication.getId()).isEqualTo(DEFAULT_COMMUNICATION_ID),
            () -> assertThat(communication.getVehicleId()).isEqualTo(DEFAULT_VEHICLE_ID),
            () -> assertThat(communication.isControlled()).isEqualTo(DEFAULT_CONTROLLED),
            () -> assertThat(communication.getCommands()).isEqualTo(newSeveralCommands(DEFAULT_GENERATOR.get())),
            () -> assertThat(communication.getTelemetry()).isEqualTo(newNormalTelemetry())
        );
    }

    /**
     * CommunicationエンティティからDTOに分解する。
     */
    @Test
    public void takeApartCommunicationTest() {
        CommunicationComponentDto dto = CommunicationFactory.takeApart(
                newSeveralCommandsCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_GENERATOR.get()
                )
        );

        assertThat(dto)
                .isEqualTo(newNormalCommunicationComponentDto(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_GENERATOR.get()
                ));
    }
}