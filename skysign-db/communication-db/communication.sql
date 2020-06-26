CREATE TABLE communication (
    id character varying(36) NOT NULL,
    mission_id character varying(36),
    version integer NOT NULL,
    CONSTRAINT communication_pkey PRIMARY KEY (id)
)