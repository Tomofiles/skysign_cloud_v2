package net.tomofiles.skysign.communication.infra.common;

import lombok.Data;

@Data
public class DeleteCondition {
    private String id;
    private int version;
}