CREATE TABLE vehicles (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    communication_id character varying(36) NOT NULL,
    is_carbon_copy boolean NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT vehicles_pkey PRIMARY KEY (id)
);
CREATE INDEX vehicles_all_select_idx ON vehicles (is_carbon_copy);
CREATE INDEX vehicles_upd_del_idx ON vehicles (id, version);