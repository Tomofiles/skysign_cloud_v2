package net.tomofiles.skysign.mission.domain.mission;

import java.util.UUID;

public class GeneratorUUID implements Generator {
    
    @Override
    public MissionId newMissionId() {
        return new MissionId(UUID.randomUUID().toString());
    }

    @Override
    public Version newVersion() {
        return new Version(UUID.randomUUID().toString());
    }
}