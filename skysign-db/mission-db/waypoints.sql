CREATE TABLE waypoints (
    mission_id character varying(36) NOT NULL,
    point_order integer NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    height_wgs84_ellipsoid_m double precision NOT NULL,
    speed_m_s double precision NOT NULL
);
CREATE INDEX waypoints_select_del_idx ON waypoints (mission_id);