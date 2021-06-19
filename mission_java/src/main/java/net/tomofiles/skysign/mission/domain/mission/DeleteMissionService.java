package net.tomofiles.skysign.mission.domain.mission;

public class DeleteMissionService {
    public static void delete(
        MissionRepository repository,
        Mission mission
    ) {
        if (mission == null) {
            return;
        }

        if (mission.isCarbonCopy()) {
            throw new CannotChangeMissionException("cannot delete carbon copied mission");
        }

        repository.remove(mission.getId(), mission.getVersion());
    }
}