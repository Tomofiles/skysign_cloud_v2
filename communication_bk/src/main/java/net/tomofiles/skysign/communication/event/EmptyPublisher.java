package net.tomofiles.skysign.communication.event;

import lombok.NoArgsConstructor;

@NoArgsConstructor
public class EmptyPublisher implements Publisher {

    @Override
    public void publish(Event event) {
    }
}