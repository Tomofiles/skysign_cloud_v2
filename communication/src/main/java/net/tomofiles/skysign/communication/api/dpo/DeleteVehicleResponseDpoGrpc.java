package net.tomofiles.skysign.communication.api.dpo;

import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.service.dpo.DeleteVehicleResponseDpo;

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