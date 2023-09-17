DROP TABLE IF EXISTS sources;
DROP TABLE IF EXISTS follow_sources;

ALTER TABLE articles
DROP INDEX sources_idx,
DROP COLUMN source_id;