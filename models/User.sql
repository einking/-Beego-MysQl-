CREATE TABLE User (
  UserId INT NOT NULL,
  Uname VARCHAR(45) NOT NULL,
  Upassword VARCHAR(45) NOT NULL,
  Confidence INT default 5,
  PRIMARY KEY (UserId));




