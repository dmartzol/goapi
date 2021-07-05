
\c api-template
SET client_min_messages TO WARNING;

CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;

CREATE TABLE accounts (
    id uuid PRIMARY KEY NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    dob date NOT NULL,
    gender CITEXT DEFAULT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE CHECK ((confirmed_email AND review_time IS NOT NULL) OR NOT active),
    email CITEXT NOT NULL UNIQUE,
    confirmed_email BOOLEAN DEFAULT FALSE,
    phone_number VARCHAR UNIQUE DEFAULT NULL,
    confirmed_phone BOOLEAN DEFAULT FALSE,
    passhash TEXT NOT NULL,
    failed_logins_count INT DEFAULT 0,
    door_code VARCHAR DEFAULT NULL,
    external_payment_customer_id INT DEFAULT NULL CHECK ((confirmed_email and review_time IS NOT NULL) OR external_payment_customer_id IS NULL),
    review_time TIMESTAMPTZ DEFAULT NULL,
    created_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY NOT NULL,
    account_id BIGINT REFERENCES accounts (id) NOT NULL,
    last_activity_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    token UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
    expiration_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 year',
    created_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;