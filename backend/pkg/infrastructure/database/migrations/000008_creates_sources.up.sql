CREATE TABLE IF NOT EXISTS sources(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(50) NULL,
  icon_url VARCHAR(100) NULL,
  domain VARCHAR(50) NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_domain (domain)
);

CREATE TABLE IF NOT EXISTS follow_sources(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  account_id INT NOT NULL,
  source_id INT NOT NULL,
  UNIQUE (account_id, source_id),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY accounts_idx (account_id),
  KEY sources_idx (source_id)
);

ALTER TABLE articles
ADD COLUMN source_id INT NOT NULL,
ADD KEY sources_idx (source_id)