BEGIN;

CREATE TABLE IF NOT EXISTS story_content (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    franchise_address VARCHAR,
    collection_address VARCHAR,
    token_id INTEGER, 
    content_json VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON story_content 
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END