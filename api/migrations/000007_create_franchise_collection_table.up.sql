BEGIN

CREATE TABLE IF NOT EXISTS franchise_collection (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    franchise_id BIGINT NOT NULL,
    collection_address VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON franchise_collection
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END