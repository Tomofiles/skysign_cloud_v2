CREATE TABLE telemetry (
    comm_id character varying(36) NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    altitude double precision NOT NULL,
    relative_altitude double precision NOT NULL,
    speed double precision NOT NULL,
    armed boolean NOT NULL,
    flight_mode character varying(50) NOT NULL,
    orientation_x double precision NOT NULL,
    orientation_y double precision NOT NULL,
    orientation_z double precision NOT NULL,
    orientation_w double precision NOT NULL,
    CONSTRAINT telemetry_pkey PRIMARY KEY (comm_id)
);