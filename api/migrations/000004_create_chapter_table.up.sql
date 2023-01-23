BEGIN;

CREATE TABLE IF NOT EXISTS chapter (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    seq_num INTEGER NOT NULL,
    title VARCHAR,
    cover_url VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON chapter
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END