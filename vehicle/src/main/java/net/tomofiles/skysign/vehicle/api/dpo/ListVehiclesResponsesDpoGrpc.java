package net.tomofiles.skysign.vehicle.api.dpo;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.service.dpo.ListVehiclesResponsesDpo;

public class ListVehiclesResponsesDpoGrpc implements ListVehiclesResponsesDpo {

    private List<Vehicle> vehicles;

    public ListVehiclesResponsesDpoGrpc() {
        this.vehicles = new ArrayList<>();
    }

    @Override
    public void setVehicles(List<Vehicle> vehicles) {
        this.vehicles = vehicles;
    }

    public proto.skysign.ListVehiclesResponses getGrpcResponse() {
        List<proto.skysign.common.Vehicle> r = this.vehicles.stream()
                .map(vehicle -> {
                    return proto.skysign.common.Vehicle.newBuilder()
                            .setId(vehicle.getId().getId())
                            .setName(vehicle.getVehicleName())
                            .setCommId(vehicle.getCommId().getId())
                            .build();
                })
                .collect(Collectors.toList());
        return proto.skysign.ListVehiclesResponses.newBuilder()
                .addAllVehicles(r)
                .build(); 
    }
}