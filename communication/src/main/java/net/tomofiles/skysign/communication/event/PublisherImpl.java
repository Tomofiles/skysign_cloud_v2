package net.tomofiles.skysign.communication.event;

import org.springframework.context.ApplicationEventPublisher;
import org.springframework.context.ApplicationEventPublisherAware;
import org.springframework.stereotype.Component;

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