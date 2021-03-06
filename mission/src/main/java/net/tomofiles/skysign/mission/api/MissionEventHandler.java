package net.tomofiles.skysign.mission.api;

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
import net.tomofiles.skysign.mission.api.event.CarbonCopyMissionRequestDpoEvent;
import net.tomofiles.skysign.mission.api.proto.MissionCopiedWhenFlightplanCopiedEventPb;
import net.tomofiles.skysign.mission.service.ManageMissionService;

@Component
@RequiredArgsConstructor
public class MissionEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(MissionEventHandler.class);
    
    private final ManageMissionService manageMissionService;

    @Value("${skysign.event.flightplan.mission_copied_when_copied_event}")
    @Setter
    private String EXCHANGE_NAME_COPIED_EVENT;

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.flightplan.mission_copied_when_copied_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.flightplan.mission_copied_when_copied_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processMissionCopiedWhenFlightplanCopiedEvent(byte[] message) throws Exception {
        MissionCopiedWhenFlightplanCopiedEventPb eventPb = new MissionCopiedWhenFlightplanCopiedEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", EXCHANGE_NAME_COPIED_EVENT, eventPb);
        CarbonCopyMissionRequestDpoEvent requestDpo = new CarbonCopyMissionRequestDpoEvent(eventPb.getEvent());
        this.manageMissionService.carbonCopyMission(requestDpo);
    }

}