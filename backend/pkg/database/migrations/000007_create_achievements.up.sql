CREATE TABLE IF NOT EXISTS achievements(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(50) NULL,
  icon_url INT NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts_to_achievements(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  achievement_id INT NOT NULL,
  account_id INT NOT NULL,
  points INT NOT NULL DEFAULT 0,
  earned_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY achievements_idx (achievement_id),
  KEY accounts_idx (account_id)
)