BEGIN;

CREATE TABLE IF NOT EXISTS character_info (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    franchise_address VARCHAR NOT NULL,
    character_id BIGINT,
    character_name VARCHAR NOT NULL,
    owner_address VARCHAR NOT NULL,
    image_url VARCHAR,
    backstory VARCHAR,
    media_uri VARCHAR, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_character_info_on_franchise_address ON character_info(franchise_address);
CREATE INDEX idx_character_info_on_character_id ON character_info(character_id);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON character_info
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END