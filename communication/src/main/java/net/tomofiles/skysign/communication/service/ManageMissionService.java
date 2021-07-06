package net.tomofiles.skysign.communication.service;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.mission.Mission;
import net.tomofiles.skysign.communication.domain.mission.MissionRepository;
import net.tomofiles.skysign.communication.service.dpo.CreateMissionRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.DeleteMissionRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.ManageMissionResponseDpo;

@Component
@AllArgsConstructor
public class ManageMissionService {

    private final MissionRepository missionRepository;

    @Transactional
    public void createMission(CreateMissionRequestDpo requestDpo, ManageMissionResponseDpo responseDpo) {
        Mission mission = new Mission(requestDpo.getMissionId());
        requestDpo.getWaypoints().stream()
            .forEach(waypoint -> {
                mission.pushWaypoint(
                    waypoint.getLatitude(), 
                    waypoint.getLongitude(), 
                    waypoint.getRelativeHeight(), 
                    waypoint.getSpeed());
            });

        this.missionRepository.save(mission);

        responseDpo.setMission(mission);
    }

    @Transactional
    public void deleteMission(DeleteMissionRequestDpo requestDpo, ManageMissionResponseDpo responseDpo) {
        Mission mission = this.missionRepository.getById(requestDpo.getMissionId());

        if (mission == null) {
            return;
        }

        this.missionRepository.remove(mission.getId());

        responseDpo.setMission(mission);
    }
}