CREATE TABLE fleets (
    id character varying(36) NOT NULL,
    flightplan_id character varying(36) NOT NULL,
    version character varying(36) NOT NULL,
    CONSTRAINT fleets_pkey PRIMARY KEY (id)
);
CREATE INDEX fleets_select_idx ON fleets (flightplan_id);
-- CREATE INDEX fleets_upd_del_idx ON fleets (id, version);