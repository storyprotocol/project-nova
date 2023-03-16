BEGIN;

ALTER TABLE story_chapter
DROP COLUMN heading,
DROP COLUMN release_at; 

END