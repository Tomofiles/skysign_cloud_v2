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
import net.tomofiles.skysign.communication.domain.communication.TelemetryUpdatedEvent;

import static net.tomofiles.skysign.communication.api.GrpcObjectMother.newNormalTelemetryGrpc;
import static net.tomofiles.skysign.communication.domain.communication.SnapshotObjectMother.newNormalTelemetrySnapshot;

public class TelemetryUpdatedEventHandlerTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final String EXCHANGE_NAME = "exchange_name";

    @Mock
    private RabbitTemplate rabbitTemplate;

    private TelemetryUpdatedEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new TelemetryUpdatedEventHandler(this.rabbitTemplate);
        this.eventHandler.setEXCHANGE_NAME(EXCHANGE_NAME);
    }

    /**
     * CommunicationのTelemetryが更新されたときに発行されるイベントを
     * 受信した場合の処理を確認する。<br>
     * RabbitMQクライアントに送信したことを検証する。
     */
    @Test
    public void fireTelemetryUpdatedEvent() throws Exception {
        TelemetryUpdatedEvent event = new TelemetryUpdatedEvent(
                DEFAULT_COMMUNICATION_ID,
                newNormalTelemetrySnapshot()
        );

        this.eventHandler.processTelemetryUpdatedEvent(event);

        ArgumentCaptor<String> exchangeCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> routingCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<Message> messageCaptor = ArgumentCaptor.forClass(Message.class);
        verify(this.rabbitTemplate, times(1)).send(
            exchangeCaptor.capture(),
            routingCaptor.capture(),
            messageCaptor.capture());

        byte[] expectMessage = proto.skysign.event.TelemetryUpdatedEvent.newBuilder()
            .setCommunicationId(DEFAULT_COMMUNICATION_ID.getId())
            .setTelemetry(newNormalTelemetryGrpc())
            .build()
            .toByteArray();

        assertThat(exchangeCaptor.getValue()).isEqualTo(EXCHANGE_NAME);
        assertThat(messageCaptor.getValue().getBody()).isEqualTo(expectMessage);
    }
}