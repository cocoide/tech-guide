CREATE TABLE IF NOT EXISTS contributions(
  account_id INT NOT NULL,
  points INT NOT NULL,
  date DATE NOT NULL,
  UNIQUE (account_id, date),
  KEY accounts_idx (account_id)
)