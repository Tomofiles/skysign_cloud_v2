CREATE TABLE navigations (
    mission_id character varying(36) NOT NULL,
    takeoff_point_ground_altitude_m double precision NOT NULL,
    upload_id character varying(36) NOT NULL,
    CONSTRAINT navigations_pkey PRIMARY KEY (mission_id)
);
