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
import net.tomofiles.skysign.vehicle.domain.vehicle.CopiedVehicleCreatedEvent;
import net.tomofiles.skysign.vehicle.infra.event.listener.proto.CopiedVehicleCreatedEventPb;

@Component
@RequiredArgsConstructor
public class CopiedVehicleCreatedEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(CopiedVehicleCreatedEventHandler.class);
    
    private final RabbitTemplate rabbitTemplate;

    @Value("${skysign.event.exchange.copied_vehicle_created_event}")
    @Setter
    private String EXCHANGE_NAME;

    @TransactionalEventListener(phase = TransactionPhase.AFTER_COMMIT)
    @Async
    public void processCopiedVehicleCreatedEvent(CopiedVehicleCreatedEvent event) {
        CopiedVehicleCreatedEventPb eventPb = new CopiedVehicleCreatedEventPb(event);
        logger.info("PUBLISH , Event: {}, Message: {}", EXCHANGE_NAME, eventPb);
        this.rabbitTemplate.send(
            EXCHANGE_NAME,
            "",
            eventPb.getMessage());
    }
}