CREATE TABLE navigations (
    mission_id character varying(36) NOT NULL,
    takeoff_point_ground_height_wgs84_ellipsoid_m double precision NOT NULL,
    CONSTRAINT navigations_pkey PRIMARY KEY (mission_id)
);
