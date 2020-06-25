package net.tomofiles.skysign.communication.infra.communication;

import lombok.Data;

@Data
public class CommunicationRecord {
    private String id;
    private String missionId;
    private int version;
}