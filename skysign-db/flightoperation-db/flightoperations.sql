CREATE TABLE flightoperations (
    id character varying(36) NOT NULL,
    flightplan_id character varying(36) NOT NULL,
    CONSTRAINT flightoperations_pkey PRIMARY KEY (id)
);