BEGIN;

CREATE TABLE IF NOT EXISTS story_info_v2 (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    franchise_id BIGINT NOT NULL,
    story_id BIGINT,
    story_name VARCHAR NOT NULL,
    story_description VARCHAR,
    owner_address VARCHAR,
    cover_url VARCHAR,
    content VARCHAR,
    media_uri VARCHAR, 
    txhash VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_story_info_v2_on_franchise_id ON story_info_v2(franchise_id);
CREATE INDEX idx_story_info_v2_on_story_id ON story_info_v2(story_id);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON story_info_v2
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END