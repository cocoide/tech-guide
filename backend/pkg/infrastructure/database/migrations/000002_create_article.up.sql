CREATE TABLE IF NOT EXISTS articles(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  original_url VARCHAR(255) NOT NULL UNIQUE,
  thumbnail_url VARCHAR(1000) NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)