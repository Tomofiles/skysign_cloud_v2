package net.tomofiles.skysign.vehicle.api.dpo;

import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.service.dpo.CreateVehicleResponseDpo;

public class CreateVehicleResponseDpoGrpc implements CreateVehicleResponseDpo {

    private Vehicle vehicle = null;

    @Override
    public void setVehicle(Vehicle vehicle) {
        this.vehicle = vehicle;
    }

    public proto.skysign.common.Vehicle getGrpcResponse() {
        return proto.skysign.common.Vehicle.newBuilder()
                .setId(vehicle.getId().getId())
                .setName(vehicle.getVehicleName())
                .setCommId(vehicle.getCommId().getId())
                .build();
    }
}