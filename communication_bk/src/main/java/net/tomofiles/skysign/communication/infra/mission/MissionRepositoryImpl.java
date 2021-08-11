package net.tomofiles.skysign.communication.infra.mission;

import java.util.List;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.mission.Mission;
import net.tomofiles.skysign.communication.domain.mission.MissionId;
import net.tomofiles.skysign.communication.domain.mission.MissionRepository;

@Component
@AllArgsConstructor
public class MissionRepositoryImpl implements MissionRepository {

    private final MissionMapper missionMapper;
    private final WaypointMapper waypointMapper;

    @Override
    public void save(Mission mission) {
        boolean isCreate = false;

        MissionRecord missionRecord = this.missionMapper.find(mission.getId().getId());

        if (missionRecord == null) {
            missionRecord = new MissionRecord();
            missionRecord.setId(mission.getId().getId());
            isCreate = true;
        }

        List<WaypointRecord> waypointRecords = mission.getWaypoints().stream()
                .map(c -> {
                        return new WaypointRecord(
                                mission.getId().getId(),
                                c.getOrder(),
                                c.getLatitude(),
                                c.getLongitude(),
                                c.getRelativeHeightM(),
                                c.getSpeedMS());
                })
                .collect(Collectors.toList());

        if (isCreate) {
            this.missionMapper.create(missionRecord);
            waypointRecords.stream()
                    .forEach(this.waypointMapper::create);
        } else {
            this.waypointMapper.delete(missionRecord.getId());
            waypointRecords.stream()
                    .forEach(this.waypointMapper::create);
        }
    }

    @Override
    public Mission getById(MissionId id) {
        MissionRecord missionRecord = this.missionMapper.find(id.getId());

        if (missionRecord == null) {
            return null;
        }

        List<WaypointRecord> waypointRecords = this.waypointMapper.find(id.getId());

        Mission mission = new Mission(id);
        waypointRecords.stream()
        .forEach(waypoint -> {
            mission.pushWaypoint(
                waypoint.getLatitudeDegree(),
                waypoint.getLongitudeDegree(),
                waypoint.getRelativeHeightM(),
                waypoint.getSpeedMS()
            );
        });

        return mission;
    }

    @Override
    public void remove(MissionId id) {
        this.missionMapper.delete(id.getId());
        this.waypointMapper.delete(id.getId());
    }
}