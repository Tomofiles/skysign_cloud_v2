package net.tomofiles.skysign.communication.domain.mission;

public interface MissionRepository {
    Mission getById(MissionId id);
    void remove(MissionId id);
    void save(Mission mission);
}