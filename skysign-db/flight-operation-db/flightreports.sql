CREATE TABLE flightreports (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    description character varying(1000) NOT NULL,
    fleet_id character varying(36) NOT NULL,
    CONSTRAINT flightreports_pkey PRIMARY KEY (id)
);