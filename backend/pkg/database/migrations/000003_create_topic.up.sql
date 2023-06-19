CREATE TABLE IF NOT EXISTS topics(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  icon_url VARCHAR(255) NULL,
  name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS articles_to_topics(
  article_id INT NOT NULL,
  topic_id INT NOT NULL,
  weight INT NOT NULL DEFAULT 3,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (article_id, topic_id),
  FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE,
  FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS follow_topics(
  topic_id INT NOT NULL,
  account_id INT NOT NULL,
  UNIQUE (account_id, topic_id),
  FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE,
  FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
)
