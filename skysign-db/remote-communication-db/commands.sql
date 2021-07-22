CREATE TABLE commands (
    id character varying(36) NOT NULL,
    communication_id character varying(36) NOT NULL,
    type character varying(50) NOT NULL,
    time TIMESTAMP NOT NULL,
    CONSTRAINT commands_pkey PRIMARY KEY (id)
);
CREATE INDEX commands_all_idx ON commands (communication_id);