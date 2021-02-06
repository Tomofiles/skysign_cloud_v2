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
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdRemovedEvent;
import net.tomofiles.skysign.communication.infra.event.listener.proto.CommunicationIdRemovedEventPb;

@Component
@RequiredArgsConstructor
public class CommunicationIdRemovedEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(CommunicationIdRemovedEventHandler.class);
    
    private final RabbitTemplate rabbitTemplate;

    @Value("${skysign.event.communication_id_removed_event}")
    @Setter
    private String EXCHANGE_NAME;

    @TransactionalEventListener(phase = TransactionPhase.AFTER_COMMIT)
    @Async
    public void processCommunicationIdRemovedEvent(CommunicationIdRemovedEvent event) {
        CommunicationIdRemovedEventPb eventPb = new CommunicationIdRemovedEventPb(event);
        logger.info("PUBLISH , Event: {}, Message: {}", EXCHANGE_NAME, eventPb);
        this.rabbitTemplate.send(
            EXCHANGE_NAME,
            "",
            eventPb.getMessage());
    }
}