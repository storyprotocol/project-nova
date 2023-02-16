BEGIN;

CREATE TABLE IF NOT EXISTS nft_collection (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    collection_address VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    symbol VARCHAR NOT NULL,
    total_cap INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON nft_collection
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END