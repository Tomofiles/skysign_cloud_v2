CREATE TABLE flightplans (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    description character varying(1000) NOT NULL,
    is_carbon_copy boolean NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT flightplans_pkey PRIMARY KEY (id)
);
CREATE INDEX flightplans_all_select_idx ON flightplans (is_carbon_copy);
-- CREATE INDEX flightplans_upd_del_idx ON flightplans (id, version);