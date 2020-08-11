package net.tomofiles.skysign.mission.infra.mission;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.domain.mission.MissionFactory;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.MissionRepository;
import net.tomofiles.skysign.mission.domain.mission.Version;
import net.tomofiles.skysign.mission.domain.mission.component.MissionComponentDto;
import net.tomofiles.skysign.mission.domain.mission.component.WaypointComponentDto;
import net.tomofiles.skysign.mission.infra.common.DeleteCondition;

@Component
@AllArgsConstructor
public class MissionRepositoryImpl implements MissionRepository {

    private final MissionMapper missionMapper;
    private final WaypointMapper waypointMapper;

    @Override
    public void save(Mission mission) {
        boolean isCreate = false;

        MissionComponentDto missionDto = MissionFactory.takeApart(mission); 

        MissionRecord missionRecord = this.missionMapper.find(missionDto.getId());

        if (missionRecord == null) {
            missionRecord = new MissionRecord();
            missionRecord.setId(missionDto.getId());
            isCreate = true;
        }

        missionRecord.setName(missionDto.getName());
        missionRecord.setTakeoffPointGroundHeightWGS84EllipsoidM(missionDto.getTakeoffPointGroundHeightWGS84M());
        missionRecord.setVersion(missionDto.getVersion());
        missionRecord.setNewVersion(missionDto.getNewVersion());

        List<WaypointRecord> waypointRecords = missionDto.getWaypoints().stream()
                .map(c -> {
                        return new WaypointRecord(
                                missionDto.getId(),
                                c.getOrder(),
                                c.getLatitude(),
                                c.getLongitude(),
                                c.getHeightWGS84M(),
                                c.getSpeedMS());
                })
                .collect(Collectors.toList());

        if (isCreate) {
            this.missionMapper.create(missionRecord);
            waypointRecords.stream()
                    .forEach(this.waypointMapper::create);
        } else {
            this.missionMapper.update(missionRecord);
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

        return MissionFactory.assembleFrom(
                new MissionComponentDto(
                        id.getId(),
                        missionRecord.getName(),
                        missionRecord.getTakeoffPointGroundHeightWGS84EllipsoidM(),
                        missionRecord.getVersion(),
                        missionRecord.getNewVersion(),
                        waypointRecords.stream()
                                .sorted(Comparator.comparing(WaypointRecord::getOrder))
                                .map(waypoint -> {
                                    return new WaypointComponentDto(
                                        waypoint.getOrder(),
                                        waypoint.getLatitude(),
                                        waypoint.getLongitude(),
                                        waypoint.getHeightWGS84EllipsoidM(),
                                        waypoint.getSpeedMS()
                                    );
                                })
                                .collect(Collectors.toList())
                )
        );
    }

    @Override
    public List<Mission> getAll() {
        List<MissionRecord> missionRecords = this.missionMapper.findAll();

        if (missionRecords.isEmpty()) {
            return Collections.emptyList();
        }

        List<Mission> missions = new ArrayList<>();
        for (MissionRecord missionRecord : missionRecords) {

            List<WaypointRecord> waypointRecords = this.waypointMapper.find(missionRecord.getId());

            Mission mission = MissionFactory.assembleFrom(
                    new MissionComponentDto(
                            missionRecord.getId(),
                            missionRecord.getName(),
                            missionRecord.getTakeoffPointGroundHeightWGS84EllipsoidM(),
                            missionRecord.getVersion(),
                            missionRecord.getNewVersion(),
                            waypointRecords.stream()
                                    .sorted(Comparator.comparing(WaypointRecord::getOrder))
                                    .map(waypoint -> {
                                        return new WaypointComponentDto(
                                            waypoint.getOrder(),
                                            waypoint.getLatitude(),
                                            waypoint.getLongitude(),
                                            waypoint.getHeightWGS84EllipsoidM(),
                                            waypoint.getSpeedMS()
                                        );
                                    })
                                    .collect(Collectors.toList())
                    )
            );

            missions.add(mission);
        }

        return missions;
    }

    @Override
    public void remove(MissionId id, Version version) {
        DeleteCondition condition = new DeleteCondition();
        condition.setId(id.getId());
        condition.setVersion(version.getVersion());

        this.missionMapper.delete(condition);
        this.waypointMapper.delete(id.getId());
    }
}