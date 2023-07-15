CREATE TABLE IF NOT EXISTS posts(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(200),
  content TEXT,
  category VARCHAR(100),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  status ENUM('publish', 'draft', 'trash')
);