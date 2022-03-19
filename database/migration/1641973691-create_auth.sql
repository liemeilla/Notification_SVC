-- Migration: create_auth
-- Created at: 2022-01-12 14:48:11
-- ====  UP  ====

BEGIN;

CREATE TABLE auth(
    customer_id integer NOT NULL,
    api_key text NOT NULL,
    status varchar(30) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    PRIMARY KEY (customer_id)
);

COMMIT;

-- ==== DOWN ====

BEGIN;

DROP TABLE auth;

COMMIT;
