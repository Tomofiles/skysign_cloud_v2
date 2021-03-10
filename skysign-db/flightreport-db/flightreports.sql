CREATE TABLE flightreports (
    id character varying(36) NOT NULL,
    flightoperation_id character varying(36) NOT NULL,
    CONSTRAINT flightreports_pkey PRIMARY KEY (id)
);