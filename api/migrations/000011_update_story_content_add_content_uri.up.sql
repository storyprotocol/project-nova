BEGIN;

ALTER TABLE story_content
ADD COLUMN content_uri VARCHAR;

END