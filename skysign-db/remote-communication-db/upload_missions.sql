CREATE TABLE upload_missions (
    id character varying(36) NOT NULL,
    communication_id character varying(36) NOT NULL,
    mission_id character varying(36) NOT NULL,
    CONSTRAINT upload_missions_pkey PRIMARY KEY (id)
);
CREATE INDEX upload_missions_all_idx ON upload_missions (communication_id);