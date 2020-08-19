package net.tomofiles.skysign.mission.generator;

import java.util.UUID;

import org.springframework.stereotype.Component;

import net.tomofiles.skysign.mission.domain.mission.Generator;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.Version;

@Component
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