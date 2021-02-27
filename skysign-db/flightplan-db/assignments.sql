CREATE TABLE assignments (
    id character varying(36) NOT NULL,
    fleet_id character varying(36) NOT NULL,
    vehicle_id character varying(36) NOT NULL,
    CONSTRAINT assignments_pkey PRIMARY KEY (id)
);
CREATE INDEX assignments_select_idx ON assignments (fleet_id);