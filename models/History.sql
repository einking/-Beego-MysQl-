CREATE TABLE History (
  CardNum VARCHAR(45) NOT NULL,
  Operate VARCHAR(45) NOT NULL,
  Timee VARCHAR(45) NOT NULL,
  FOREIGN KEY (CardNum) REFERENCES Card(CardNum) on delete cascade on update cascade);