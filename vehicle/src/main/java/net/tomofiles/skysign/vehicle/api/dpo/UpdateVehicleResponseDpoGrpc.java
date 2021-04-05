package net.tomofiles.skysign.vehicle.api.dpo;

import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.service.dpo.UpdateVehicleResponseDpo;

public class UpdateVehicleResponseDpoGrpc implements UpdateVehicleResponseDpo {

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