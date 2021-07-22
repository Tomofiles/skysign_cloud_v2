CREATE TABLE fleets (
    id character varying(36) NOT NULL,
    is_carbon_copy boolean NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT fleets_pkey PRIMARY KEY (id)
);
