-- Migration: create_notification_log
-- Created at: 2022-01-12 14:45:05
-- ====  UP  ====

BEGIN;

CREATE TABLE notification_log(
    id text NOT NULL,
    customer_id integer NOT NULL,
    notification_url text NOT NULL,
    notification_data json NOT NULL,
    request_json json NOT NULL,
    response_json json NOT NULL,
    status_sent varchar(30) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);

COMMIT;

-- ==== DOWN ====

BEGIN;

DROP TABLE notification_log;

COMMIT;
