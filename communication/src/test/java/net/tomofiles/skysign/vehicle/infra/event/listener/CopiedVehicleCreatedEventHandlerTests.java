package net.tomofiles.skysign.vehicle.infra.event.listener;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.core.RabbitTemplate;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;

import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationId;
import net.tomofiles.skysign.vehicle.domain.vehicle.CopiedVehicleCreatedEvent;
import net.tomofiles.skysign.vehicle.domain.vehicle.FlightplanId;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;

public class CopiedVehicleCreatedEventHandlerTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final FlightplanId DEFAULT_FLIGHTPLAN_ID = new FlightplanId(UUID.randomUUID().toString());
    private static final String EXCHANGE_NAME = "exchange_name";

    @Mock
    private RabbitTemplate rabbitTemplate;

    private CopiedVehicleCreatedEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new CopiedVehicleCreatedEventHandler(this.rabbitTemplate);
        this.eventHandler.setEXCHANGE_NAME(EXCHANGE_NAME);
    }

    /**
     * コピーされたVehicleが作成されたときに発行されるイベントを
     * 受信した場合の処理を確認する。<br>
     * RabbitMQクライアントに送信したことを検証する。
     */
    @Test
    public void fireCopiedVehicleCreatedEvent() throws Exception {
        CopiedVehicleCreatedEvent event = new CopiedVehicleCreatedEvent(
                DEFAULT_VEHICLE_ID,
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_FLIGHTPLAN_ID
        );

        this.eventHandler.processCopiedVehicleCreatedEvent(event);

        ArgumentCaptor<String> exchangeCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> routingCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<Message> messageCaptor = ArgumentCaptor.forClass(Message.class);
        verify(this.rabbitTemplate, times(1)).send(
            exchangeCaptor.capture(),
            routingCaptor.capture(),
            messageCaptor.capture());

        byte[] expectMessage = proto.skysign.event.CopiedVehicleCreatedEvent.newBuilder()
            .setVehicleId(DEFAULT_VEHICLE_ID.getId())
            .setCommunicationId(DEFAULT_COMMUNICATION_ID.getId())
            .setFlightplanId(DEFAULT_FLIGHTPLAN_ID.getId())
            .build()
            .toByteArray();

        assertThat(exchangeCaptor.getValue()).isEqualTo(EXCHANGE_NAME);
        assertThat(messageCaptor.getValue().getBody()).isEqualTo(expectMessage);
    }
}