package net.tomofiles.skysign.communication.domain.communication;

public interface CommunicationRepository {
    Communication getById(CommunicationId id);
    void remove(CommunicationId id);
    void save(Communication communication);
}