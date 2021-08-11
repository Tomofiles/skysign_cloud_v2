package net.tomofiles.skysign.communication.domain.communication;

import java.util.List;

public interface CommunicationRepository {
    List<Communication> getAll();
    Communication getById(CommunicationId id);
    void remove(CommunicationId id);
    void save(Communication communication);
}