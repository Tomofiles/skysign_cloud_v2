CREATE TABLE flightoperations (
    id character varying(36) NOT NULL,
    flightplan_id character varying(36) NOT NULL,
    is_completed boolean NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT flightoperations_pkey PRIMARY KEY (id)
);
CREATE INDEX flightoperations_all_select_idx ON flightoperations (is_completed);