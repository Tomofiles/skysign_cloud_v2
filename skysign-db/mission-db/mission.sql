CREATE TABLE mission (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    takeoff_point_ground_height_wgs84_ellipsoid_m double precision NOT NULL,
    is_carbon_copy boolean NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT mission_pkey PRIMARY KEY (id)
);
CREATE INDEX mission_all_select_idx ON mission (is_carbon_copy);
CREATE INDEX mission_select_upd_del_idx ON mission (id, version);