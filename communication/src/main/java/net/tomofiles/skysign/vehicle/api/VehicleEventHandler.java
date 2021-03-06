package net.tomofiles.skysign.vehicle.api;

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
import net.tomofiles.skysign.vehicle.api.event.CarbonCopyVehicleRequestDpoEvent;
import net.tomofiles.skysign.vehicle.api.proto.VehicleCopiedWhenFlightplanCopiedEventPb;
import net.tomofiles.skysign.vehicle.service.ManageVehicleService;

@Component
@RequiredArgsConstructor
public class VehicleEventHandler {
    private static final Logger logger = LoggerFactory.getLogger(VehicleEventHandler.class);
    
    private final ManageVehicleService manageVehicleService;

    @Value("${skysign.event.flightplan.vehicle_copied_when_copied_event}")
    @Setter
    private String EXCHANGE_NAME_COPIED_EVENT;

    @RabbitListener(
        bindings = @QueueBinding(
            value = @Queue(value = "${skysign.event.flightplan.vehicle_copied_when_copied_event}", durable = "false", exclusive = "false", autoDelete = "true"),
            exchange = @Exchange(value = "${skysign.event.flightplan.vehicle_copied_when_copied_event}", type = "fanout", durable = "false", autoDelete = "true")
        )
    )
    public void processVehicleCopiedWhenFlightplanCopiedEvent(byte[] message) throws Exception {
        VehicleCopiedWhenFlightplanCopiedEventPb eventPb = new VehicleCopiedWhenFlightplanCopiedEventPb(message);
        logger.info("RECEIVE , Event: {}, Message: {}", EXCHANGE_NAME_COPIED_EVENT, eventPb);
        CarbonCopyVehicleRequestDpoEvent requestDpo = new CarbonCopyVehicleRequestDpoEvent(eventPb.getEvent());
        this.manageVehicleService.carbonCopyVehicle(requestDpo);
    }

}