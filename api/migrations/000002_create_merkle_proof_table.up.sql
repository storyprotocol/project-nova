BEGIN;

CREATE TABLE IF NOT EXISTS merkle_proof (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    proof VARCHAR NOT NULL,
    allowlist_id uuid NOT NULL,
    wallet_address uuid NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_merkle_proof_on_allowlist_id ON merkle_proof(allowlist_id);
CREATE INDEX idx_merkle_proof_on_wallet_address ON merkle_proof(wallet_address);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON merkle_proof
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END