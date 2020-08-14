package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.usecase.dpo.CreateVehicleRequestDpo;
import proto.skysign.CreateVehicleRequest;

@RequiredArgsConstructor
public class CreateVehicleRequestDpoGrpc implements CreateVehicleRequestDpo {

    private final CreateVehicleRequest request;

    @Override
    public String getVehicleName() {
        return this.request.getName();
    }

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getCommId());
    }
}