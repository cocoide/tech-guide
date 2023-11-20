CREATE TABLE job_settings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    worker_id INT NOT NULL,
    interval_seconds INT NOT NULL,
    scheduled_at DATETIME NULL,
    status INT NOT NULL,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE job_settings
    ADD INDEX worker_idx (worker_id),
    ADD INDEX status_idx (status),
    ADD INDEX scheduled_at_idx (scheduled_at);