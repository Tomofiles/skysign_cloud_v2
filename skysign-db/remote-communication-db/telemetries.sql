CREATE TABLE telemetries (
    communication_id character varying(36) NOT NULL,
    latitude_degree double precision NOT NULL,
    longitude_degree double precision NOT NULL,
    altitude_m double precision NOT NULL,
    relative_altitude_m double precision NOT NULL,
    speed_ms double precision NOT NULL,
    armed boolean NOT NULL,
    flight_mode character varying(50) NOT NULL,
    orientation_x double precision NOT NULL,
    orientation_y double precision NOT NULL,
    orientation_z double precision NOT NULL,
    orientation_w double precision NOT NULL,
    CONSTRAINT telemetries_pkey PRIMARY KEY (communication_id)
);