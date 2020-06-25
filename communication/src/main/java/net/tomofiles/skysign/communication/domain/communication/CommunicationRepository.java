package net.tomofiles.skysign.communication.domain.communication;

import net.tomofiles.skysign.communication.domain.common.Version;

public interface CommunicationRepository {
    Communication getById(CommunicationId id);
    void remove(CommunicationId id, Version version);
    void save(Communication communication);
}