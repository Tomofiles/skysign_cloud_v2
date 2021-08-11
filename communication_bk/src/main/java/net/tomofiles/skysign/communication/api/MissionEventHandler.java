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
import net.tomofiles.skysign.communication.api.proto.CopiedMissionCreatedEventPb;
import net.tomofiles.skysign.communication.service.ManageMissionService;

@Component
@RequiredArgsConstructor
public class MissionEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(MissionEventHandler.class);
    
    private final ManageMissionService manageMissionService;

    @Value("${skysign.event.queue.copied_mission_created_event}")
    @Setter
    private String QUEUE_NAME_CREATED_EVENT;

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.queue.copied_mission_created_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.exchange.copied_mission_created_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processCopiedMissionCreatedEvent(byte[] message) throws Exception {
        CopiedMissionCreatedEventPb eventPb = new CopiedMissionCreatedEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", QUEUE_NAME_CREATED_EVENT, eventPb);
        CreateMissionRequestDpoEvent requestDpo = new CreateMissionRequestDpoEvent(eventPb.getEvent());
        this.manageMissionService.createMission(requestDpo, mission -> {/** 何もしない */});
    }
}