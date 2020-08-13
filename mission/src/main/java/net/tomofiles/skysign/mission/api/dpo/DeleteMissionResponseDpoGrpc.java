package net.tomofiles.skysign.mission.api.dpo;

import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.service.dpo.DeleteMissionResponseDpo;

public class DeleteMissionResponseDpoGrpc implements DeleteMissionResponseDpo {

    private Mission mission = null;

    @Override
    public void setMission(Mission mission) {
        this.mission = mission;
    }

    public boolean isEmpty() {
        return mission == null;
    }
}