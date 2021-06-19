package net.tomofiles.skysign.mission.service;

import java.util.List;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.mission.domain.mission.CarbonCopyMissionService;
import net.tomofiles.skysign.mission.domain.mission.DeleteMissionService;
import net.tomofiles.skysign.mission.domain.mission.Generator;
import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.domain.mission.MissionFactory;
import net.tomofiles.skysign.mission.domain.mission.MissionRepository;
import net.tomofiles.skysign.mission.service.dpo.CarbonCopyMissionRequestDpo;
import net.tomofiles.skysign.mission.service.dpo.CreateMissionRequestDpo;
import net.tomofiles.skysign.mission.service.dpo.CreateMissionResponseDpo;
import net.tomofiles.skysign.mission.service.dpo.DeleteMissionRequestDpo;
import net.tomofiles.skysign.mission.service.dpo.DeleteMissionResponseDpo;
import net.tomofiles.skysign.mission.service.dpo.GetMissionRequestDpo;
import net.tomofiles.skysign.mission.service.dpo.GetMissionResponseDpo;
import net.tomofiles.skysign.mission.service.dpo.ListMissionsResponsesDpo;
import net.tomofiles.skysign.mission.service.dpo.UpdateMissionRequestDpo;
import net.tomofiles.skysign.mission.service.dpo.UpdateMissionResponseDpo;

@Component
@AllArgsConstructor
public class ManageMissionService {

    private final MissionRepository missionRepository;
    private final Generator generator;

    @Transactional
    public void listMissions(ListMissionsResponsesDpo responsesDpo) {
        List<Mission> missions = this.missionRepository.getAllOriginal();

        responsesDpo.setMissions(missions);
    }

    @Transactional
    public void getMission(GetMissionRequestDpo requestDpo, GetMissionResponseDpo responseDpo) {
        Mission mission = this.missionRepository.getById(requestDpo.getMissionId());

        responseDpo.setMission(mission);
    }

    @Transactional
    public void createMission(CreateMissionRequestDpo requestDpo, CreateMissionResponseDpo responseDpo) {
        Mission mission = MissionFactory.newInstance(this.generator);

        mission.nameMission(requestDpo.getMissionName());
        mission.replaceNavigationWith(requestDpo.getNavigation());

        this.missionRepository.save(mission);

        responseDpo.setMission(mission);
    }

    @Transactional
    public void updateMission(UpdateMissionRequestDpo requestDpo, UpdateMissionResponseDpo responseDpo) {
        Mission mission = this.missionRepository.getById(requestDpo.getMissionId());

        if (mission == null) {
            return;
        }

        mission.nameMission(requestDpo.getMissionName());
        mission.replaceNavigationWith(requestDpo.getNavigation());

        this.missionRepository.save(mission);

        responseDpo.setMission(mission);
    }

    @Transactional
    public void deleteMission(DeleteMissionRequestDpo requestDpo, DeleteMissionResponseDpo responseDpo) {
        Mission mission = this.missionRepository.getById(requestDpo.getMissionId());
        
        DeleteMissionService.delete(this.missionRepository, mission);

        responseDpo.setMission(mission);
    }

    @Transactional
    public void carbonCopyMission(CarbonCopyMissionRequestDpo requestDpo) {
        CarbonCopyMissionService.copy(
            this.generator, 
            this.missionRepository, 
            requestDpo.getOriginalId(), 
            requestDpo.getNewId());
    }
}