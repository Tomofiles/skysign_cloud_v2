package net.tomofiles.skysign.communication.event;

public interface Publisher {
    void publish(final Event event);
}