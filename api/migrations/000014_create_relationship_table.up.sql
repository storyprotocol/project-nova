BEGIN;

CREATE TABLE IF NOT EXISTS relationships (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    source_contract VARCHAR NOT NULL,
    dest_contract VARCHAR NOT NULL,
    src_id BIGINT NOT NULL,
    dst_id BIGINT NOT NULL,
		type VARCHAR NOT NULL,
    tx_hash VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_relationships_on_src_id ON relationships(src_id);
CREATE INDEX idx_relationships_on_dst_id ON relationships(dst_id);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON relationships
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

END;
