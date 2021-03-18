package net.tomofiles.skysign.vehicle.event;

import lombok.NoArgsConstructor;

@NoArgsConstructor
public class EmptyPublisher implements Publisher {

    @Override
    public void publish(Event event) {
    }
}