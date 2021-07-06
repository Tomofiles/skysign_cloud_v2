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
import net.tomofiles.skysign.communication.api.event.CreateMissionRequestDpoEvent;
import net.tomofiles.skysign.communication.api.event.DeleteMissionRequestDpoEvent;
import net.tomofiles.skysign.communication.api.proto.MissionCreatedEventPb;
import net.tomofiles.skysign.communication.api.proto.MissionDeletedEventPb;
import net.tomofiles.skysign.communication.service.ManageMissionService;

@Component
@RequiredArgsConstructor
public class MissionEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(MissionEventHandler.class);
    
    private final ManageMissionService manageMissionService;

    @Value("${skysign.event.queue.mission_created_event}")
    @Setter
    private String QUEUE_NAME_CREATED_EVENT;

    @Value("${skysign.event.queue.mission_deleted_event}")
    @Setter
    private String QUEUE_NAME_DELETED_EVENT;

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.queue.mission_created_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.exchange.mission_created_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processMissionCreatedEvent(byte[] message) throws Exception {
        MissionCreatedEventPb eventPb = new MissionCreatedEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", QUEUE_NAME_CREATED_EVENT, eventPb);
        CreateMissionRequestDpoEvent requestDpo = new CreateMissionRequestDpoEvent(eventPb.getEvent());
        this.manageMissionService.createMission(requestDpo, mission -> {/** 何もしない */});
    }

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.queue.mission_deleted_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.exchange.mission_deleted_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processMissionDeletedEvent(byte[] message) throws Exception {
        MissionDeletedEventPb eventPb = new MissionDeletedEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", QUEUE_NAME_DELETED_EVENT, eventPb);
        DeleteMissionRequestDpoEvent requestDpo = new DeleteMissionRequestDpoEvent(eventPb.getEvent());
        this.manageMissionService.deleteMission(requestDpo, mission -> {/** 何もしない */});
    }
}