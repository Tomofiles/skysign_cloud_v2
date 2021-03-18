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
import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationIdRemovedEvent;
import net.tomofiles.skysign.vehicle.domain.vehicle.Version;

public class CommunicationIdRemovedEventHandlerTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final String EXCHANGE_NAME = "exchange_name";

    @Mock
    private RabbitTemplate rabbitTemplate;

    private CommunicationIdRemovedEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new CommunicationIdRemovedEventHandler(this.rabbitTemplate);
        this.eventHandler.setEXCHANGE_NAME(EXCHANGE_NAME);
    }

    /**
     * Vehicleが更新されたときにCommunicationIDが削除されたイベントを
     * 受信した場合の処理を確認する。<br>
     * RabbitMQクライアントに送信したことを検証する。
     */
    @Test
    public void fireCommunicationIdGaveEvent() throws Exception {
        CommunicationIdRemovedEvent event = new CommunicationIdRemovedEvent(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_VERSION
        );

        this.eventHandler.processCommunicationIdRemovedEvent(event);

        ArgumentCaptor<String> exchangeCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> routingCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<Message> messageCaptor = ArgumentCaptor.forClass(Message.class);
        verify(this.rabbitTemplate, times(1)).send(
            exchangeCaptor.capture(),
            routingCaptor.capture(),
            messageCaptor.capture());

        byte[] expectMessage = proto.skysign.event.CommunicationIdRemovedEvent.newBuilder()
            .setCommunicationId(DEFAULT_COMMUNICATION_ID.getId())
            .setVersion(DEFAULT_VERSION.getVersion())
            .build()
            .toByteArray();

        assertThat(exchangeCaptor.getValue()).isEqualTo(EXCHANGE_NAME);
        assertThat(messageCaptor.getValue().getBody()).isEqualTo(expectMessage);
    }
}