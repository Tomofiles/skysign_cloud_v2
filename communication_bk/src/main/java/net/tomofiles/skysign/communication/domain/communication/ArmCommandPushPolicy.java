package net.tomofiles.skysign.communication.domain.communication;

public class ArmCommandPushPolicy {
    public static boolean isFollow(CommandType commandType, Communication communication) {
        if (commandType == CommandType.TAKEOFF || commandType == CommandType.START) {
            if (!communication.getTelemetry().isArmed()) {
                return true;
            }
        }
        return false;
    }
}