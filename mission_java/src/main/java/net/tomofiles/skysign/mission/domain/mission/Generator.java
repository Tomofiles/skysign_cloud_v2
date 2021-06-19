package net.tomofiles.skysign.mission.domain.mission;

public interface Generator {
    public MissionId newMissionId();
    public Version newVersion();
}