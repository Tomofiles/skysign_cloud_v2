package net.tomofiles.skysign.mission.domain.mission;

public class CarbonCopyMissionService {
    public static void copy(
        Generator generator,
        MissionRepository repository,
        MissionId originalId,
        MissionId newId
    ) {
        Mission newMission = repository.getById(newId);

        if (newMission != null) {
            return;
        }

        Mission original = repository.getById(originalId);

        if (original == null) {
            return;
        }

        Mission copy = MissionFactory.copy(original, newId, generator);

        repository.save(copy);

    }
}