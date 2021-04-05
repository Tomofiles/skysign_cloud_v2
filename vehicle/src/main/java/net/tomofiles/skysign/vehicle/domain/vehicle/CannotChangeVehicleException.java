package net.tomofiles.skysign.vehicle.domain.vehicle;

public class CannotChangeVehicleException extends RuntimeException {

    private static final long serialVersionUID = 1L;
    
    public CannotChangeVehicleException(String arg0) {
        super(arg0);
    }

    public CannotChangeVehicleException(Throwable cause) {
        super(cause);
    }

    public CannotChangeVehicleException() {
        super();
    }

    @Override
    public String toString() {
        return super.toString();
    }
}