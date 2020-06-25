package net.tomofiles.skysign.communication.api;

import java.time.LocalTime;
import java.util.List;
import java.util.stream.Collectors;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.infra.VehicleRepository;
import proto.skysign.CreateVehicleRequest;
import proto.skysign.ListVehiclesRequest;
import proto.skysign.ListVehiclesResponses;
import proto.skysign.Vehicle;
import proto.skysign.VehicleServiceGrpc.VehicleServiceImplBase;

@GRpcService
@Controller
public class Endpoint extends VehicleServiceImplBase {
    
    @Autowired
    private VehicleRepository vehicleRepository;

    @Override
    public void listVehicles(ListVehiclesRequest request, StreamObserver<ListVehiclesResponses> responseObserver) {
        System.out.println("list");

        List<net.tomofiles.skysign.communication.infra.Vehicle> all = vehicleRepository.findAll();

        List<Vehicle> r = all.stream()
            .map(v -> Vehicle.newBuilder().setId(v.getId()).setName(v.getName()).build())
            .collect(Collectors.toList());
        
        responseObserver.onNext(ListVehiclesResponses.newBuilder().addAllVehicles(r).build()); 
        responseObserver.onCompleted();
    }

    @Override
    public void createVehicle(CreateVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
        System.out.println("create");
        net.tomofiles.skysign.communication.infra.Vehicle v = new net.tomofiles.skysign.communication.infra.Vehicle();
        v.setId(LocalTime.now().toString());
        v.setName("vehicle");
        
        vehicleRepository.create(v);
        
        Vehicle r = Vehicle.newBuilder().setId(v.getId()).setName(v.getName()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}