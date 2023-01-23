BEGIN;

CREATE TABLE IF NOT EXISTS nft (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    collection_address VARCHAR NOT NULL,
    token_id INTEGER NOT NULL,
    owner_address VARCHAR,
    image_url VARCHAR,
    traits VARCHAR,
    backstory VARCHAR(3000),
    owner_updated_at TIMESTAMP NOT NULL DEFAULT now(),
    story_updated_at TIMESTAMP
);

CREATE INDEX idx_nft_on_collection_address ON nft(collection_address);
CREATE INDEX idx_nft_on_owner_address ON nft(owner_address);

END