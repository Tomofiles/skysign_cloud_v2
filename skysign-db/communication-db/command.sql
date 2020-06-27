CREATE TABLE command (
    id character varying(36) NOT NULL,
    comm_id character varying(36) NOT NULL,
    type character varying(50) NOT NULL,
    CONSTRAINT command_pkey PRIMARY KEY (id)
);
CREATE INDEX command_all_idx ON command (comm_id);