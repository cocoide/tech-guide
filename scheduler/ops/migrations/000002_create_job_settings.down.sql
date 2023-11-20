DROP INDEX worker_idx ON job_settings;
DROP INDEX status_idx ON job_settings;
DROP INDEX scheduled_at_idx ON job_settings;

DROP TABLE IF EXISTS job_settings;