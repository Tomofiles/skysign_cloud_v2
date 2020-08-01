CREATE TABLE waypoint (
    mission_id character varying(36) NOT NULL,
    order integer NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    height_wgs84_ellipsoid_m double precision NOT NULL,
    speed_m_s double precision NOT NULL,
    CONSTRAINT waypoint_pkey PRIMARY KEY (mission_id)
);
CREATE INDEX waypoint_select_del_idx ON mission (mission_id);