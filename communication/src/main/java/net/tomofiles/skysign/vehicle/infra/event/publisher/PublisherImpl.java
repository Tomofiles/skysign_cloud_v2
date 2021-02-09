package net.tomofiles.skysign.vehicle.infra.event.publisher;

import org.springframework.context.ApplicationEventPublisher;
import org.springframework.context.ApplicationEventPublisherAware;
import org.springframework.stereotype.Component;

import net.tomofiles.skysign.vehicle.event.Event;
import net.tomofiles.skysign.vehicle.event.Publisher;

@Component
public class PublisherImpl implements ApplicationEventPublisherAware, Publisher {

    private ApplicationEventPublisher publisher;

    @Override
    public void setApplicationEventPublisher(ApplicationEventPublisher applicationEventPublisher) {
        this.publisher = applicationEventPublisher;
    }

    @Override
    public void publish(Event event) {
        this.publisher.publishEvent(event);
    }
}