package net.tomofiles.skysign.communication.api;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.amqp.rabbit.annotation.QueueBinding;
import org.springframework.amqp.rabbit.annotation.Queue;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.amqp.rabbit.annotation.Exchange;
import org.springframework.stereotype.Component;

import lombok.AllArgsConstructor;
import lombok.RequiredArgsConstructor;
import lombok.Setter;
import net.tomofiles.skysign.communication.api.dpo.CreateCommunicationRequestDpoEvent;
import net.tomofiles.skysign.communication.api.dpo.RecreateCommunicationRequestDpoEvent;
import net.tomofiles.skysign.communication.api.proto.CommunicationIdChangedEventPb;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;
import net.tomofiles.skysign.communication.service.ManageCommunicationService;

@Component
@RequiredArgsConstructor
public class CommunicationEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(CommunicationEventHandler.class);
    
    private final ManageCommunicationService manageCommunicationService;

    @Value("${skysign.event.communication_id_changed_event}")
    @Setter
    private String EXCHANGE_NAME;

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.communication_id_changed_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.communication_id_changed_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processCommunicationIdChangedEvent(byte[] message) throws Exception {
        CommunicationIdChangedEventPb eventPb = new CommunicationIdChangedEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", EXCHANGE_NAME, eventPb);
        CommunicationIdChangedEvent event = eventPb.getEvent();
        if (event.isFirst()) {
            CreateCommunicationRequestDpoEvent requestDpo = new CreateCommunicationRequestDpoEvent(event);
            this.manageCommunicationService.createCommunication(requestDpo, communication -> {/** 何もしない */});
        } else {
            RecreateCommunicationRequestDpoEvent requestDpo = new RecreateCommunicationRequestDpoEvent(event);
            this.manageCommunicationService.recreateCommunication(requestDpo, communication -> {/** 何もしない */});
        }
    }
}