package net.tomofiles.skysign.communication.infra.event.listener;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Component;
import org.springframework.transaction.event.TransactionPhase;
import org.springframework.transaction.event.TransactionalEventListener;

import lombok.RequiredArgsConstructor;
import lombok.Setter;
import net.tomofiles.skysign.communication.domain.communication.TelemetryUpdatedEvent;
import net.tomofiles.skysign.communication.infra.event.listener.proto.TelemetryUpdatedEventPb;

@Component
@RequiredArgsConstructor
public class TelemetryUpdatedEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(TelemetryUpdatedEventHandler.class);
    
    private final RabbitTemplate rabbitTemplate;

    @Value("${skysign.event.telemetry_updated_event}")
    @Setter
    private String EXCHANGE_NAME;

    @TransactionalEventListener(phase = TransactionPhase.AFTER_COMMIT)
    @Async
    public void processTelemetryUpdatedEvent(TelemetryUpdatedEvent event) {
        TelemetryUpdatedEventPb eventPb = new TelemetryUpdatedEventPb(event);
        logger.info("PUBLISH , Event: {}, Message: {}", EXCHANGE_NAME, eventPb);
        this.rabbitTemplate.send(
            EXCHANGE_NAME,
            "",
            eventPb.getMessage());
    }
}