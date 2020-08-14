package net.tomofiles.skysign.communication.usecase.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public interface GetCommunicationRequestDpo {
    public CommunicationId getCommId();
}