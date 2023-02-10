BEGIN;

CREATE TABLE IF NOT EXISTS story_info (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    franchise_id BIGINT NOT NULL,
    seq_num INTEGER NOT NULL,
    title VARCHAR NOT NULL,
    subtitle VARCHAR,
    cover_url VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(franchise_id, seq_num)
);

CREATE INDEX idx_story_info_on_franchise_id ON story_info(franchise_id);
CREATE INDEX idx_story_info_on_seq_num ON story_info(seq_num);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON story_info
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END