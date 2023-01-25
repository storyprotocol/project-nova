BEGIN;

CREATE TABLE IF NOT EXISTS wallet_merkle_proof (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    proof VARCHAR NOT NULL,
    allowlist_id uuid NOT NULL,
    wallet_address VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_wallet_merkle_proof_on_allowlist_id ON wallet_merkle_proof(allowlist_id);
CREATE INDEX idx_wallet_merkle_proof_on_wallet_address ON wallet_merkle_proof(wallet_address);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON wallet_merkle_proof
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END