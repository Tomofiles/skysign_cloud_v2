package net.tomofiles.skysign.vehicle.infra.event.listener;

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
import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationIdGaveEvent;
import net.tomofiles.skysign.vehicle.infra.event.listener.proto.CommunicationIdGaveEventPb;

@Component
@RequiredArgsConstructor
public class CommunicationIdGaveEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(CommunicationIdGaveEventHandler.class);
    
    private final RabbitTemplate rabbitTemplate;

    @Value("${skysign.event.communication_id_gave_event}")
    @Setter
    private String EXCHANGE_NAME;

    @TransactionalEventListener(phase = TransactionPhase.AFTER_COMMIT)
    @Async
    public void processCommunicationIdGaveEvent(CommunicationIdGaveEvent event) {
        CommunicationIdGaveEventPb eventPb = new CommunicationIdGaveEventPb(event);
        logger.info("PUBLISH , Event: {}, Message: {}", EXCHANGE_NAME, eventPb);
        this.rabbitTemplate.send(
            EXCHANGE_NAME,
            "",
            eventPb.getMessage());
    }
}