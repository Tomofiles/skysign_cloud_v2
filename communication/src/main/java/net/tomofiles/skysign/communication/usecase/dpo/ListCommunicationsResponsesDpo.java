package net.tomofiles.skysign.communication.usecase.dpo;

import java.util.List;

import net.tomofiles.skysign.communication.domain.communication.Communication;

public interface ListCommunicationsResponsesDpo {
    public void setCommunications(List<Communication> communications);
}