CREATE TABLE missions (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    is_carbon_copy boolean NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT missions_pkey PRIMARY KEY (id)
);
CREATE INDEX missions_all_select_idx ON missions (is_carbon_copy);
CREATE INDEX missions_select_upd_del_idx ON missions (id, version);
