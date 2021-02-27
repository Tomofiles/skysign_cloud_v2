package net.tomofiles.skysign.vehicle.api.dpo;

import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.service.dpo.DeleteVehicleResponseDpo;

public class DeleteVehicleResponseDpoGrpc implements DeleteVehicleResponseDpo {

    private Vehicle vehicle = null;

    @Override
    public void setVehicle(Vehicle vehicle) {
        this.vehicle = vehicle;
    }

    public boolean isEmpty() {
        return this.vehicle == null;
    }
}