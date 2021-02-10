CREATE TABLE upload_mission (
    id character varying(36) NOT NULL,
    comm_id character varying(36) NOT NULL,
    mission_id character varying(36) NOT NULL,
    CONSTRAINT upload_mission_pkey PRIMARY KEY (id)
);
CREATE INDEX upload_mission_all_idx ON upload_mission (comm_id);