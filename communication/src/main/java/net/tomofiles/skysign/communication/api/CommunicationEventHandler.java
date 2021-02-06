package net.tomofiles.skysign.communication.api;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.amqp.rabbit.annotation.QueueBinding;
import org.springframework.amqp.rabbit.annotation.Queue;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.amqp.rabbit.annotation.Exchange;
import org.springframework.stereotype.Component;

import lombok.RequiredArgsConstructor;
import lombok.Setter;
import net.tomofiles.skysign.communication.api.dpo.CreateCommunicationRequestDpoEvent;
import net.tomofiles.skysign.communication.api.dpo.DeleteCommunicationRequestDpoEvent;
import net.tomofiles.skysign.communication.api.proto.CommunicationIdGaveEventPb;
import net.tomofiles.skysign.communication.api.proto.CommunicationIdRemovedEventPb;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdGaveEvent;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdRemovedEvent;
import net.tomofiles.skysign.communication.service.ManageCommunicationService;

@Component
@RequiredArgsConstructor
public class CommunicationEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(CommunicationEventHandler.class);
    
    private final ManageCommunicationService manageCommunicationService;

    @Value("${skysign.event.communication_id_gave_event}")
    @Setter
    private String EXCHANGE_NAME_GAVE_EVENT;

    @Value("${skysign.event.communication_id_removed_event}")
    @Setter
    private String EXCHANGE_NAME_REMOVED_EVENT;

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.communication_id_gave_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.communication_id_gave_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processCommunicationIdGaveEvent(byte[] message) throws Exception {
        CommunicationIdGaveEventPb eventPb = new CommunicationIdGaveEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", EXCHANGE_NAME_GAVE_EVENT, eventPb);
        CommunicationIdGaveEvent event = eventPb.getEvent();
        CreateCommunicationRequestDpoEvent requestDpo = new CreateCommunicationRequestDpoEvent(event);
        this.manageCommunicationService.createCommunication(requestDpo, communication -> {/** 何もしない */});
    }

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.communication_id_removed_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.communication_id_removed_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processCommunicationIdRemovedEvent(byte[] message) throws Exception {
        CommunicationIdRemovedEventPb eventPb = new CommunicationIdRemovedEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", EXCHANGE_NAME_REMOVED_EVENT, eventPb);
        CommunicationIdRemovedEvent event = eventPb.getEvent();
        DeleteCommunicationRequestDpoEvent requestDpo = new DeleteCommunicationRequestDpoEvent(event);
        this.manageCommunicationService.deleteCommunication(requestDpo, communication -> {/** 何もしない */});
    }
}