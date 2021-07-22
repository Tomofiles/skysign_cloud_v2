CREATE TABLE events (
    id character varying(36) NOT NULL,
    fleet_id character varying(36) NOT NULL,
    assignment_id character varying(36) NOT NULL,
    mission_id character varying(36) NOT NULL,
    CONSTRAINT events_pkey PRIMARY KEY (id)
);
CREATE INDEX events_select_idx ON events (fleet_id);