package net.tomofiles.skysign.communication.infra.event.publisher;

import org.springframework.context.ApplicationEventPublisher;
import org.springframework.context.ApplicationEventPublisherAware;
import org.springframework.stereotype.Component;

import net.tomofiles.skysign.communication.event.Event;
import net.tomofiles.skysign.communication.event.Publisher;

@Component("communication-publisher-impl")
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