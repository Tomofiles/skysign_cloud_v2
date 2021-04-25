CREATE TABLE actions (
    id character varying(36) NOT NULL,
    communication_id character varying(36) NOT NULL,
    flightplan_id character varying(36) NOT NULL,
    is_completed boolean NOT NULL,
    CONSTRAINT actions_pkey PRIMARY KEY (id)
);
