BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS whitelist_wallet (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    address VARCHAR(255),
    merkle_proof VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

END