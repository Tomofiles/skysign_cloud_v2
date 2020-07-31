package net.tomofiles.skysign.mission.domain.mission;

import java.util.List;

public interface MissionRepository {
    void save(Mission mission);
    Mission getById(MissionId id);
    List<Mission> getAll();
    void remove(MissionId id, Version version);
}