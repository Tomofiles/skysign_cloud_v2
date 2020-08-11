package net.tomofiles.skysign.mission.api;

import java.util.List;
import java.util.NoSuchElementException;
import java.util.stream.Collectors;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;
import org.springframework.transaction.annotation.Transactional;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.mission.domain.mission.Generator;
import net.tomofiles.skysign.mission.domain.mission.GeodesicCoordinates;
import net.tomofiles.skysign.mission.domain.mission.Height;
import net.tomofiles.skysign.mission.domain.mission.MissionFactory;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.MissionRepository;
import net.tomofiles.skysign.mission.domain.mission.Navigation;
import net.tomofiles.skysign.mission.domain.mission.Speed;
import proto.skysign.DeleteMissionRequest;
import proto.skysign.Empty;
import proto.skysign.GetMissionRequest;
import proto.skysign.ListMissionsRequest;
import proto.skysign.ListMissionsResponses;
import proto.skysign.Mission;
import proto.skysign.MissionItem;
import proto.skysign.ManageMissionServiceGrpc.ManageMissionServiceImplBase;

@GRpcService
@Controller
@AllArgsConstructor
public class ManageMissionEndpoint extends ManageMissionServiceImplBase {

    private final MissionRepository missionRepository;
    private final Generator generator;

    @Override
    @Transactional
    public void listMissions(ListMissionsRequest request, StreamObserver<ListMissionsResponses> responseObserver) {
        List<net.tomofiles.skysign.mission.domain.mission.Mission> missions;

        try {
            missions = missionRepository.getAll();
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        List<Mission> r = missions.stream().map(mission -> {
            return Mission.newBuilder()
                    .setId(mission.getId().getId())
                    .setName(mission.getMissionName())
                    .setTakeoffPointGroundHeight(mission.getNavigation().getTakeoffPointGroundHeight().getHeightM())
                    .addAllItems(mission.getNavigation().getWaypoints().stream().map(waypoint -> {
                        return MissionItem.newBuilder()
                                .setLatitude(waypoint.getLatitude())
                                .setLongitude(waypoint.getLongitude())
                                .setRelativeHeight(waypoint.getRelativeHeightM())
                                .setSpeed(waypoint.getSpeedMS())
                                .build();
                        }).collect(Collectors.toList()))
                    .build();
        }).collect(Collectors.toList());

        responseObserver.onNext(ListMissionsResponses.newBuilder().addAllMissions(r).build());
        responseObserver.onCompleted();
    }

    @Override
    @Transactional
    public void getMission(GetMissionRequest request, StreamObserver<Mission> responseObserver) {
        MissionId id = new MissionId(request.getId());
        net.tomofiles.skysign.mission.domain.mission.Mission mission;

        try {
            mission = missionRepository.getById(id);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (mission == null) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("mission-idに合致するMissionが存在しません。")
                    .asRuntimeException());
            return;
        }

        Mission r = Mission.newBuilder()
                .setId(mission.getId().getId())
                .setName(mission.getMissionName())
                .setTakeoffPointGroundHeight(mission.getNavigation().getTakeoffPointGroundHeight().getHeightM())
                .addAllItems(mission.getNavigation().getWaypoints().stream().map(waypoint -> {
                    return MissionItem.newBuilder()
                            .setLatitude(waypoint.getLatitude())
                            .setLongitude(waypoint.getLongitude())
                            .setRelativeHeight(waypoint.getRelativeHeightM())
                            .setSpeed(waypoint.getSpeedMS())
                            .build();
                    }).collect(Collectors.toList()))
                .build();
        responseObserver.onNext(r);
        responseObserver.onCompleted();
    }

    @Override
    @Transactional
    public void createMission(Mission request, StreamObserver<Mission> responseObserver) {
        net.tomofiles.skysign.mission.domain.mission.Mission mission = MissionFactory.newInstance(this.generator);

        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(request.getTakeoffPointGroundHeight()));
        request.getItemsList()
                .forEach(item -> {
                    navigation.pushNextWaypoint(
                        new GeodesicCoordinates(item.getLatitude(), item.getLongitude()),
                        Height.fromM(item.getRelativeHeight()),
                        Speed.fromMS(item.getSpeed()));
                });

        mission.nameMission(request.getName());
        mission.replaceNavigationWith(navigation);

        try {
            this.missionRepository.save(mission);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        Mission r = Mission.newBuilder(request)
                .setId(mission.getId().getId())
                .build();
        responseObserver.onNext(r);
        responseObserver.onCompleted();
    }

    @Override
    @Transactional
    public void updateMission(Mission request, StreamObserver<Mission> responseObserver) {
        MissionId id = new MissionId(request.getId());
        net.tomofiles.skysign.mission.domain.mission.Mission mission;

        try {
            mission = missionRepository.getById(id);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (mission == null) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("mission-idに合致するMissionが存在しません。")
                    .asRuntimeException());
            return;
        }

        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(request.getTakeoffPointGroundHeight()));
        request.getItemsList()
                .forEach(item -> {
                    navigation.pushNextWaypoint(
                        new GeodesicCoordinates(item.getLatitude(), item.getLongitude()),
                        Height.fromM(item.getRelativeHeight()),
                        Speed.fromMS(item.getSpeed()));
                });

        mission.nameMission(request.getName());
        mission.replaceNavigationWith(navigation);

        try {
            this.missionRepository.save(mission);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        Mission r = Mission.newBuilder(request)
                .setId(mission.getId().getId())
                .build();
        responseObserver.onNext(r);
        responseObserver.onCompleted();
    }

    @Override
    @Transactional
    public void deleteMission(DeleteMissionRequest request, StreamObserver<Empty> responseObserver) {
        MissionId id = new MissionId(request.getId());
        net.tomofiles.skysign.mission.domain.mission.Mission mission;

        try {
            mission = missionRepository.getById(id);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (mission == null) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("mission-idに合致するMissionが存在しません。")
                    .asRuntimeException());
            return;
        }

        try {
            this.missionRepository.remove(id, mission.getVersion());
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(Empty.newBuilder().build()); 
        responseObserver.onCompleted();
    }
}