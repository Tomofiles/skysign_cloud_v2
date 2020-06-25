package net.tomofiles.skysign.communication.infra.communication;

import lombok.Data;

@Data
public class DeleteCondition {
    private String id;
    private int version;
}