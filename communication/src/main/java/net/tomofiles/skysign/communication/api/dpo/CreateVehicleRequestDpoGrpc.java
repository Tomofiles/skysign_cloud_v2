package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.CreateVehicleRequestDpo;
import proto.skysign.common.Vehicle;

@RequiredArgsConstructor
public class CreateVehicleRequestDpoGrpc implements CreateVehicleRequestDpo {

    private final Vehicle request;

    @Override
    public String getVehicleName() {
        return this.request.getName();
    }

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getCommId());
    }
}