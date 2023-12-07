BEGIN;

CREATE TABLE IF NOT EXISTS wallet_sign_info (
    wallet_address VARCHAR PRIMARY KEY,
    nonce VARCHAR,
    signature VARCHAR,
    nonce_created_at TIMESTAMP,
    signature_verified_at TIMESTAMPTZ
);

END