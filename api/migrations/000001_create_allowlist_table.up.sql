BEGIN;

CREATE OR REPLACE FUNCTION trigger_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS nft_allowlist (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    type VARCHAR NOT NULL,
    collection_address VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON nft_allowlist
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END