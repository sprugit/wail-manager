CREATE IF NOT EXISTS TABLE NODE(
  ID INT PRIMARY KEY,
  ACCOUNT VARCHAR(50),
  DOMAIN VARCHAR(50),
  PASSWORD VARCHAR(50),
  PARENT INT,
  FOREIGN KEY (PARENT) REFERENCES NODE(ID) ON DELETE CASCADE
);