package net.tomofiles.skysign.communication.domain.communication;

import java.util.ArrayList;

import net.tomofiles.skysign.communication.domain.common.Version;

public class CommunicationFactory {

    public static Communication newInstance(CommunicationId id) {
        Communication communication = new Communication(id, new ArrayList<>());
        communication.setVersion(new Version(1));
        return communication;
    }

    public static Communication rebuild(CommunicationId id, String missionId, int version) {
        Communication communication = new Communication(id, new ArrayList<>());
        communication.setMissionId(new MissionId(missionId));
        communication.setVersion(new Version(version));
        return communication;
    }
}