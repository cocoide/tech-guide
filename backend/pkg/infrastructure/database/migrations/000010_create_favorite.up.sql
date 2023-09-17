CREATE TABLE IF NOT EXISTS favorite_articles(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  account_id INT NOT NULL,
  article_id INT NOT NULL,
  UNIQUE (account_id, article_id),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY accounts_idx (account_id),
  KEY articles_idx (article_id)
);
