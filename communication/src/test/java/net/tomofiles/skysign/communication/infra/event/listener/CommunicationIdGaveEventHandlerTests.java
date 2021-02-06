package net.tomofiles.skysign.communication.infra.event.listener;

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

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdGaveEvent;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.Version;

public class CommunicationIdGaveEventHandlerTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final String EXCHANGE_NAME = "exchange_name";

    @Mock
    private RabbitTemplate rabbitTemplate;

    private CommunicationIdGaveEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new CommunicationIdGaveEventHandler(this.rabbitTemplate);
        this.eventHandler.setEXCHANGE_NAME(EXCHANGE_NAME);
    }

    /**
     * Vehicleが作成されたときにCommunicationIDが付与されたイベントを
     * 受信した場合の処理を確認する。<br>
     * RabbitMQクライアントに送信したことを検証する。
     */
    @Test
    public void fireCommunicationIdGaveEvent() throws Exception {
        CommunicationIdGaveEvent event = new CommunicationIdGaveEvent(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_VEHICLE_ID,
                DEFAULT_VERSION
        );

        this.eventHandler.processCommunicationIdGaveEvent(event);

        ArgumentCaptor<String> exchangeCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> routingCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<Message> messageCaptor = ArgumentCaptor.forClass(Message.class);
        verify(this.rabbitTemplate, times(1)).send(
            exchangeCaptor.capture(),
            routingCaptor.capture(),
            messageCaptor.capture());

        byte[] expectMessage = proto.skysign.event.CommunicationIdGaveEvent.newBuilder()
            .setCommunicationId(DEFAULT_COMMUNICATION_ID.getId())
            .setVehicleId(DEFAULT_VEHICLE_ID.getId())
            .setVersion(DEFAULT_VERSION.getVersion())
            .build()
            .toByteArray();

        assertThat(exchangeCaptor.getValue()).isEqualTo(EXCHANGE_NAME);
        assertThat(messageCaptor.getValue().getBody()).isEqualTo(expectMessage);
    }
}