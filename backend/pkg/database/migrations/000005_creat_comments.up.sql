CREATE TABLE IF NOT EXISTS comments(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  content VARCHAR(500) NULL,
  account_id INT NOT NULL,
  article_id INT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY articles_idx (article_id),
  KEY accounts_idx (account_id)
);