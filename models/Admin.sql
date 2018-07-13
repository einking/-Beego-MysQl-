CREATE TABLE Admin (
  AdminId VARCHAR(45) NOT NULL,
  Aname VARCHAR(45) NOT NULL,
  Apassword VARCHAR(45) NOT NULL,
  PRIMARY KEY (AdminId));

INSERT INTO Admin (AdminId,Aname,Apassword) VALUES
     ( '1', 'Aa', 'red'),
     ( '2', 'Bb', 'pink'),
     ( '3', 'Cc', 'yellow');
