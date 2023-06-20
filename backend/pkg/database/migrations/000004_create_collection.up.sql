CREATE TABLE IF NOT EXISTS collections(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(200) NULL,
  visibility INT NOT NULL DEFAULT 0,
  account_id INT NOT NULL,
  UNIQUE (name, account_id),
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY accounts_idx (account_id)
);

CREATE TABLE IF NOT EXISTS bookmarks(
  article_id INT NOT NULL,
  collection_id INT NOT NULL,
  UNIQUE (article_id, collection_id),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY articles_idx (article_id),
  KEY collections_idx (collection_id)
)