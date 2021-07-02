CREATE TABLE waypoints (
    mission_id character varying(36) NOT NULL,
    point_order integer NOT NULL,
    latitude_degree double precision NOT NULL,
    longitude_degree double precision NOT NULL,
    relative_height_m double precision NOT NULL,
    speed_ms double precision NOT NULL
);
CREATE INDEX waypoints_select_del_idx ON waypoints (mission_id);
