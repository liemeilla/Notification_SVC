-- Migration: create_customer_notification
-- Created at: 2022-01-12 13:51:59
-- ====  UP  ====

BEGIN;

CREATE TABLE customer_notification(
    customer_id INTEGER NOT NULL,
    notification_url text NOT NULL,
    status varchar(30) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    PRIMARY KEY (customer_id)
);

COMMIT;

-- ==== DOWN ====

BEGIN;

DROP TABLE customer_notification;

COMMIT;
