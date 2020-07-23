CREATE TABLE vehicle (
    id character varying(36) NOT NULL,
    name character varying(200) NOT NULL,
    comm_id character varying(36) NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT vehicle_pkey PRIMARY KEY (id)
);
CREATE INDEX vehicle_upd_del_idx ON vehicle (id, version);