CREATE TABLE vehicle (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    communication_id character varying(36) NOT NULL,
    is_carbon_copy boolean NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT vehicle_pkey PRIMARY KEY (id)
);
CREATE INDEX vehicle_all_select_idx ON vehicle (is_carbon_copy);
CREATE INDEX vehicle_upd_del_idx ON vehicle (id, version);