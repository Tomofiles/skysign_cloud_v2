CREATE TABLE communication (
    id character varying(36) NOT NULL,
    mission_id character varying(36),
    CONSTRAINT communication_pkey PRIMARY KEY (id)
);
CREATE INDEX communication_upd_del_idx ON communication (id);