package net.tomofiles.skysign.mission.domain.mission;

public class CannotChangeMissionException extends RuntimeException {

    private static final long serialVersionUID = 1L;
    
    public CannotChangeMissionException(String arg0) {
        super(arg0);
    }

    public CannotChangeMissionException(Throwable cause) {
        super(cause);
    }

    public CannotChangeMissionException() {
        super();
    }

    @Override
    public String toString() {
        return super.toString();
    }
}