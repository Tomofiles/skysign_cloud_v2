package net.tomofiles.skysign.communication.api.dpo;

import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.service.dpo.GetVehicleResponseDpo;

public class GetVehicleResponseDpoGrpc implements GetVehicleResponseDpo {

    private Vehicle vehicle = null;

    @Override
    public void setVehicle(Vehicle vehicle) {
        this.vehicle = vehicle;
    }

    public boolean isEmpty() {
        return this.vehicle == null;
    }

    public proto.skysign.common.Vehicle getGrpcResponse() {
        return proto.skysign.common.Vehicle.newBuilder()
                .setId(vehicle.getId().getId())
                .setName(vehicle.getVehicleName())
                .setCommId(vehicle.getCommId().getId())
                .build();
    }
}