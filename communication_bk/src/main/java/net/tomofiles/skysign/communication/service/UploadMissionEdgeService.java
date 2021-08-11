package net.tomofiles.skysign.communication.service;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.mission.Mission;
import net.tomofiles.skysign.communication.domain.mission.MissionRepository;
import net.tomofiles.skysign.communication.service.dpo.GetUploadMissionRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.GetUploadMissionResponseDpo;

@Component
@AllArgsConstructor
public class UploadMissionEdgeService {

    private final MissionRepository missionRepository;

    @Transactional
    public void getUploadMission(GetUploadMissionRequestDpo requestDpo, GetUploadMissionResponseDpo responseDpo) {
        Mission mission = this.missionRepository.getById(requestDpo.getMissionId());

        if (mission == null) {
            return;
        }

        responseDpo.setMission(mission);
    }
}