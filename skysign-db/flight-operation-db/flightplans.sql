CREATE TABLE flightplans (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    description character varying(1000) NOT NULL,
    fleet_id character varying(36) NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT flightplans_pkey PRIMARY KEY (id)
);
